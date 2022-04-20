package logic

import (
	"github.com/gin-gonic/gin"

	"integral/model"

	"sync"
)

// @Author: Feng
// @Date: 2022/3/25 15:43

type integralHandler interface {
	Modify(ctx *gin.Context, req *model.ModifyReq, rsp *model.ModifyRsp) error
	Query(ctx *gin.Context, req *model.QueryReq, rsp *model.QueryRsp) error
	QueryFlow(ctx *gin.Context, req *model.QueryFlowReq, rsp *model.QueryFlowRsp) error
	Rollback(ctx *gin.Context, req *model.RollbackReq, rsp *model.RollbackRsp) error
}

var (
	integralHandlerMap      = map[string]integralHandler{}
	integralHandlerMapMutex = sync.RWMutex{}
)

//RegisterIntegralHandler 注册积分处理器
func RegisterIntegralHandler(appid string, h integralHandler) {
	integralHandlerMapMutex.Lock()
	defer integralHandlerMapMutex.Unlock()
	integralHandlerMap[appid] = h
}

//GetIntegralHandler 获取积分处理器
func GetIntegralHandler(appid string) integralHandler {
	h, exist := integralHandlerMap[appid]
	if !exist {
		return nil
	}
	return h
}
