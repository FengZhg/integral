package dbhandler

import (
	"github.com/gin-gonic/gin"
	"integral/server"
)

// @Author: Feng
// @Date: 2022/4/11 20:21

//Rollback Redis处理器回滚
func (d *dbHandler) Rollback(ctx *gin.Context, req *server.RollbackReq, rsp *server.RollbackRsp) error {

	return nil
}
