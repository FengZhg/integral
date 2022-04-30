package redishandler

import (
	"context"
	"encoding/json"
	"github.com/FengZhg/go_tools/errs"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"integral/dao"
	"integral/dao/pulsarClient"
	"integral/model"
	"time"
)

// @Author: Feng
// @Date: 2022/3/26 14:14

//Rollback Redis处理器回滚
func (r *RedisHandler) Rollback(ctx *gin.Context, req *model.RollbackReq, rsp *model.RollbackRsp) error {
	// 进行回滚获取流水
	flowStr, err := doRollback(ctx, req)
	if err != nil {
		log.Errorf("Do Rollback Error %v", err)
		return err
	}

	// 构造并发送流水
	go doRollbackFlow(flowStr)
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
		-- redis.call('DEL', KEYS[3])
		return {flow ,0}
	`
)

//doRollback 操作进行回滚
func doRollback(ctx *gin.Context, req *model.RollbackReq) (string, error) {
	// 构造请求参数
	keys := []string{
		getBalanceKey(req.GetAppid(), req.GetType(), req.GetUid()),
		getOrderKey(req.GetAppid(), req.GetType(), req.GetUid(), req.GetOid()),
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
	code, okCode := interSli[1].(int64)
	if !okFlow || !okCode {
		log.Errorf("Parse Rollback Lua Return Error")
		return "", model.ReturnFormatError
	}

	err = errs.Code2Error(int32(code))
	if err != nil {
		log.Errorf("lua Return Error %v", err)
		return "", err
	}

	return flow, nil
}

//buildRollbackFlow 构造流水payload
func buildRollbackFlow(flowStr string) ([]byte, error) {
	// 反序列化flow
	flow := model.SingleFlow{}
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

	// 序列化流水
	flowBytes, err := json.Marshal(flow)
	if err != nil {
		log.Errorf("Marshal Flow Error %v", err)
		return nil, err
	}
	return flowBytes, nil
}

//doRollbackFlow 进行回滚
func doRollbackFlow(flowStr string) error {
	// 构建回滚结构
	flowBytes, err := buildRollbackFlow(flowStr)
	if err != nil {
		log.Errorf("Build Flow Error %v", err)
		return err
	}

	// 回滚丢进pulsar
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = pulsarClient.PulsarCfg.Produce(ctx, flowBytes)
	if err != nil {
		log.Errorf("Produce Message Error %v", err)
		return err
	}
	return nil
}
