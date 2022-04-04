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
// @Date: 2022/3/25 17:46

//Modify Redis处理器的修改函数
func (r *redisHandler) Modify(ctx *gin.Context, req *server.ModifyReq, rsp *server.ModifyRsp) error {
	// 构造流水字符串
	flowByte, err := json.Marshal(req)
	if err != nil {
		log.Errorf("Build Flow Error %v", err)
		return err
	}

	// 余额修改
	balance, err := modifyBalance(ctx, req, string(flowByte))
	if err != nil {
		log.Errorf("Modify Redis Error %v", err)
	}

	// 发送生产pulsar消息
	err = pulsarClient.Send(model.PulsarOpt, flowByte)
	if err != nil {
		log.Errorf("Pulsar Send Msg Error %v", err)
		return err
	}

	rsp.Integral = balance
	return nil
}

// 定义余额修改lua脚本
const (
	modifyScript = `
	if tonumber(redis.call('EXISTS', KEYS[2])) == 1 then 
		return {"", 10004}
	end
	local flow = ARGV[2]
	local absBalance = tonumber(ARGV[1]) or 0
	local balance = tonumber(redis.call('GET', KEYS[1])) or 0
	if balance + absBalance < 0 then 
		return {0,10002}
	end
	redis.call('SETEX', KEYS[2], 2419200, absBalance)
	redis.call('SETEX', KEYS[3], 2419200, flow)
	return {tonumber(redis.call('INCRBY',KEYS[1], absBalance)), 0}
`
)

//modifyBalance 余额修改
func modifyBalance(ctx *gin.Context, req *server.ModifyReq, flow string) (int64, error) {
	// 构造key和参数
	keys := []string{
		getOrderKey(req.GetAppid(), req.GetType(), req.GetUid(), req.GetOid()),
		getBalanceKey(req.GetAppid(), req.GetType(), req.GetUid()),
		getFlowKey(req.GetAppid(), req.GetType(), req.GetUid(), req.GetOid()),
	}
	integral := req.GetIntegral()
	if req.GetOpt() == 2 {
		integral = -1 * integral
	}
	param := []interface{}{integral, flow}

	// 请求
	rsp, err := dao.GetRedisClient().Eval(ctx, modifyScript, keys, param...).Int64Slice()
	if err != nil {
		log.Errorf("Modify Balance Redis Error %v", err)
		return 0, err
	}
	if len(rsp) != 2 {
		return 0, model.ReturnFormatError
	}

	// 解析lua脚本返回值
	balance, errCode := rsp[0], rsp[1]
	err = errs.Code2Error(int32(errCode))
	if err != nil {
		log.Errorf("lua Return Error %v", err)
		return 0, err
	}

	return balance, nil
}
