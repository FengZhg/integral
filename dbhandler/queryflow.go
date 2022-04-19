package dbhandler

import (
	"github.com/gin-gonic/gin"
	"integral/logic"
	"integral/redishandler"
)

// @Author: Feng
// @Date: 2022/4/11 20:20

//QueryFlow Redis处理器查询积分流水
func (d *dbHandler) QueryFlow(ctx *gin.Context, req *logic.QueryFlowReq, rsp *logic.QueryFlowRsp) error {
	return (&redishandler.RedisHandler{}).QueryFlow(ctx, req, rsp)
}
