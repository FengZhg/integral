package model

import "github.com/FengZhg/go_tools/errs"

// @Author: Feng
// @Date: 2022/3/25 17:18

const (
	errorParam               = 10000
	errorHandler             = 10001
	errorBalanceInsufficient = 10002
	errorFlow                = 10003
	errorModifyRepeated      = 10004
)

var (
	ParamError               = errs.NewError(errorParam, "参数错误")
	HandlerError             = errs.NewError(errorHandler, "该appid尚未注册")
	BalanceInsufficientError = errs.NewError(errorBalanceInsufficient, "余额不足")
	FlowError                = errs.NewError(errorFlow, "流水缺失")
	ModifyRepeatedError      = errs.NewError(errorModifyRepeated, "余额修改操作重复")
)
