package logic

import (
	"fmt"
	"github.com/FengZhg/go_tools/goJwt"
	"github.com/gin-gonic/gin"
	"integral/model"
)

// @Author: Feng
// @Date: 2022/3/25 15:43

//Modify 逻辑函数
func Modify(ctx *gin.Context, req *model.ModifyReq, rsp *model.ModifyRsp) error {

	// 参数校验
	err := checkModify(ctx, req)
	if err != nil {
		return err
	}

	// 获取对应处理器
	h := GetIntegralHandler(req.GetAppid())
	if h != nil {
		return h.Modify(ctx, req, rsp)
	}
	return model.HandlerError
}

//checkModify 参数校验
func checkModify(ctx *gin.Context, req *model.ModifyReq) error {
	// 校验
	if goJwt.GetLoginInfo(ctx).GetUid() != req.GetUid() {
		fmt.Println(goJwt.GetLoginInfo(ctx).String(), req.String())
		return model.ParamError
	}
	// 判空
	if req.GetType() == "" || req.GetAppid() == "" || req.GetIntegral() == 0 || req.GetOid() == "" ||
		req.GetOpt() == 0 || req.GetUid() == "" {
		return model.ParamError
	}
	return nil
}
