package logic

import (
	"github.com/FengZhg/go_tools/goJwt"
	"github.com/gin-gonic/gin"
	"integral/model"
)

// @Author: Feng
// @Date: 2022/3/25 15:46

//QueryFlow 查询积分流水
func QueryFlow(ctx *gin.Context, req *QueryFlowReq, rsp *QueryFlowRsp) error {

	// 参数校验
	if err := checkQueryFlow(ctx, req); err != nil {
		return err
	}

	// 获取对应处理器
	h := GetIntegralHandler(req.GetAppid())
	if h != nil {
		return h.QueryFlow(ctx, req, rsp)
	}
	return model.HandlerError
}

//checkQueryFlow 参数校验
func checkQueryFlow(ctx *gin.Context, req *QueryFlowReq) error {

	loginInfo := goJwt.GetLoginInfo(ctx)
	if loginInfo.GetUid() != req.GetUid() || req.GetType() == "" || req.GetAppid() == "" || req.GetNum() < 0 {
		return model.ParamError
	}

	if req.GetNum() > 20 {
		req.Num = 20
	}
	return nil
}
