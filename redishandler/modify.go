package redishandler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	flow := string(flowByte)

	// 余额修改

	return nil
}

// 定义余额修改lua脚本
const (
	modifyScript = `
	if tonumber(redis.call('EXISTS', KEYS[1])) == 1 then 
		return {"", 10004}
	end
	local flow = ARGV[2]
	if flow == nil then
		return {0,10003}
	end
	local absBalance = tonumber(ARGV[1])
	local balance = tonumber(redis.call('GET', KEYS[1])) or 0
	if balance + absBalance < 0 then 
		return {0,10002}
	end
	redis.call('INCRBY',KEY[4], absBalance)
	redis.call('SETEX', KEYS[1], 2419200, flow)
	return {balance+absBalance, 0}
`
)

//modifyBalance 余额修改
func modifyBalance(ctx *gin.Context, req *server.ModifyReq, flow string) (int64, error) {

	return 0, nil

}
