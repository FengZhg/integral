package redishandler

import (
	"github.com/gin-gonic/gin"
	"integral/server"
)

// @Author: Feng
// @Date: 2022/3/26 14:14

//Rollback Redis处理器回滚
func (r *redisHandler) Rollback(ctx *gin.Context, req *server.RollbackReq, rsp *server.RollbackRsp) error {

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

func rollback() {

}
