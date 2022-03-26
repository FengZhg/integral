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
