package model

import "github.com/FengZhg/go_tools/errs"

// @Author: Feng
// @Date: 2022/3/25 17:18

const (
	errorParam               = 10000
	errorHandler             = 10001
	errorBalanceInsufficient = 10002
	errorAlreadyRollback     = 10003
	errorModifyRepeated      = 10004
	errorReturnFormat        = 10005
	errorOrderNotExist       = 10006
)

var (
	ParamError               = errs.NewError(errorParam, "参数错误")
	HandlerError             = errs.NewError(errorHandler, "该appid尚未注册")
	BalanceInsufficientError = errs.NewError(errorBalanceInsufficient, "余额不足")
	AlreadyRollbackError     = errs.NewError(errorAlreadyRollback, "该订单已被回滚")
	ModifyRepeatedError      = errs.NewError(errorModifyRepeated, "余额已修改，操作重复")
	ReturnFormatError        = errs.NewError(errorReturnFormat, "Redis返回格式错误")
	OrderNotExistError       = errs.NewError(errorOrderNotExist, "订单不存在")
)
