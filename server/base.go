package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"integral/logic"
	"net/http"
)

// @Author: Feng
// @Date: 2022/3/25 15:28

//queryBase 积分查询基本函数
func queryBase(ctx *gin.Context) {
	// pb反序列化填充req
	req, rsp := &logic.QueryReq{}, &logic.QueryRsp{}
	err := ctx.ShouldBind(req)
	if err != nil {
		log.Errorf("Should Bind Req Error interface = Query Request = %v err = %v", ctx.Request, err)
		ctx.Error(err)
		return
	}

	// 调用handler函数
	err = logic.Query(ctx, req, rsp)
	if err != nil {
		ctx.Error(err)
		return
	}

	// 构造返回体
	ctx.JSON(http.StatusOK, rsp)
}

//modifyBase 积分修改基本函数
func modifyBase(ctx *gin.Context) {
	// pb反序列化填充req
	req, rsp := &logic.ModifyReq{}, &logic.ModifyRsp{}
	err := ctx.ShouldBind(req)
	if err != nil {
		log.Errorf("Should Bind Req Error interface = Modify Request = %v err = %v", ctx.Request, err)
		ctx.Error(err)
		return
	}

	// 调用handler函数
	err = logic.Modify(ctx, req, rsp)
	if err != nil {
		ctx.Error(err)
		return
	}

	// 构造返回体
	ctx.JSON(http.StatusOK, rsp)
}

//rollbackBase 积分修改回滚基本函数
func rollbackBase(ctx *gin.Context) {
	// pb反序列化填充req
	req, rsp := &logic.RollbackReq{}, &logic.RollbackRsp{}
	err := ctx.ShouldBind(req)
	if err != nil {
		log.Errorf("Should Bind Req Error interface = Rollback Request = %v err = %v", ctx.Request, err)
		ctx.Error(err)
		return
	}

	// 调用handler函数
	err = logic.Rollback(ctx, req, rsp)
	if err != nil {
		ctx.Error(err)
		return
	}

	// 构造返回体
	ctx.JSON(http.StatusOK, rsp)
}

//queryFlowBase 查询积分流水基本函数
func queryFlowBase(ctx *gin.Context) {
	// pb反序列化填充req
	req, rsp := &logic.QueryFlowReq{}, &logic.QueryFlowRsp{}
	err := ctx.ShouldBind(req)
	if err != nil {
		log.Errorf("Should Bind Req Error interface = QueryFlow Request = %v err = %v", ctx.Request, err)
		ctx.Error(err)
		return
	}

	// 调用handler函数
	err = logic.QueryFlow(ctx, req, rsp)
	if err != nil {
		ctx.Error(err)
		return
	}

	// 构造返回体
	ctx.JSON(http.StatusOK, rsp)
}
