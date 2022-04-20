package logic

import (
	"github.com/FengZhg/go_tools/goJwt"
	"github.com/gin-gonic/gin"
	"integral/model"
)

// @Author: Feng
// @Date: 2022/3/25 15:44

//Rollback 回滚流水
func Rollback(ctx *gin.Context, req *model.RollbackReq, rsp *model.RollbackRsp) error {
	//参数校验
	if err := checkRollback(ctx, req); err != nil {
		return err
	}

	// 获取对应处理器
	h := GetIntegralHandler(req.GetAppid())
	if h != nil {
		return h.Rollback(ctx, req, rsp)
	}
	return model.HandlerError
}

//checkRollback 参数校验
func checkRollback(ctx *gin.Context, req *model.RollbackReq) error {
	if goJwt.GetLoginInfo(ctx).GetUid() != req.GetUid() || req.GetOid() == "" || req.GetType() == "" || req.GetAppid() == "" {
		return model.ParamError
	}
	return nil
}
