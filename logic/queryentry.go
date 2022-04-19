package logic

import (
	"github.com/gin-gonic/gin"
	"integral/model"
)

// @Author: Feng
// @Date: 2022/3/25 15:34

//Query 积分查询
func Query(ctx *gin.Context, req *QueryReq, rsp *QueryRsp) error {

	// 参数校验
	if err := checkQuery(req); err != nil {
		return err
	}

	// 获取对应处理器
	h := GetIntegralHandler(req.GetAppid())
	if h != nil {
		return h.Query(ctx, req, rsp)
	}
	return model.HandlerError
}

//checkQuery 参数校验
func checkQuery(req *QueryReq) error {
	if req.GetType() == "" || req.GetAppid() == "" || len(req.GetUids()) == 0 {
		return model.ParamError
	}
	return nil
}
