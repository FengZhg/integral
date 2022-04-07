package redishandler

import (
	"encoding/json"
	"github.com/FengZhg/go_tools/errs"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"integral/dao"
	"integral/dao/pulsarClient"
	"integral/model"
	"integral/server"
)

// @Author: Feng
// @Date: 2022/3/26 14:14

//Rollback Redis处理器回滚
func (r *redisHandler) Rollback(ctx *gin.Context, req *server.RollbackReq, rsp *server.RollbackRsp) error {
	// 进行回滚获取流水
	flowStr, err := doRollback(ctx, req)
	if err != nil {
		log.Errorf("Do Rollback Error %v", err)
		return err
	}

	// 构建回滚结构
	flowBytes, err := buildFlowBytes(flowStr)
	if err != nil {
		log.Errorf("Build Flow Error %v", err)
		return err
	}

	// 回滚丢进pulsar
	err = pulsarClient.Send(model.PulsarOpt, flowBytes)
	if err != nil {
		log.Errorf("Pulsar Send Error %v", err)
		return err
	}
	return nil
}

const (
	rollbackScript = `
	local flow = redis.call('GET', KEYS[3])
	if flow == nil then
		return {'', 10003}
	end
	local absBalance = tonumber(redis.call('GET', KEYS[2]))
	if absBalance == nil then
		return {'', 10006}
	end
	redis.call('INCRBY',KEYS[1], -1 * absBalance)
	redis.call('DEL', KEYS[3])
	return {flow ,0}
`
)

//doRollback 操作进行回滚
func doRollback(ctx *gin.Context, req *server.RollbackReq) (string, error) {
	// 构造请求参数
	keys := []string{
		getOrderKey(req.GetAppid(), req.GetType(), req.GetUid(), req.GetOid()),
		getBalanceKey(req.GetAppid(), req.GetType(), req.GetUid()),
		getFlowKey(req.GetAppid(), req.GetType(), req.GetUid(), req.GetOid()),
	}

	// 发送lua脚本
	interSli, err := dao.GetRedisClient().Eval(ctx, rollbackScript, keys).Slice()
	if err != nil {
		log.Errorf("Rollback Order Redis Error %v", err)
		return "", err
	}
	if len(interSli) != 2 {
		return "", model.ReturnFormatError
	}

	// 结果assert
	flow, okFlow := interSli[0].(string)
	code, okCode := interSli[1].(int32)
	if !okFlow || !okCode {
		log.Errorf("Parse Rollback Lua Return Error")
		return "", model.ReturnFormatError
	}

	err = errs.Code2Error(code)
	if err != nil {
		log.Errorf("lua Return Error %v", err)
		return "", err
	}

	return flow, nil
}

//buildFlowBytes 构造流水payload
func buildFlowBytes(flowStr string) ([]byte, error) {
	// 反序列化flow
	flow := server.SingleFlow{}
	err := json.Unmarshal([]byte(flowStr), &flow)
	if err != nil {
		log.Errorf("Unmarshal Flow from Pulsar Message Error %v", err)
		return nil, err
	}

	// 流水修改
	if flow.GetOpt() == 1 {
		flow.Opt = 2
	} else if flow.GetOpt() == 2 {
		flow.Opt = 1
	} else {
		return nil, model.ParamError
	}
	flow.Desc = "订单回滚：" + flow.Desc

	// 序列化流水
	flowBytes, err := json.Marshal(flow)
	if err != nil {
		log.Errorf("Marshal Flow Error %v", err)
		return nil, err
	}
	return flowBytes, nil
}
