package model

import "github.com/FengZhg/go_tools/errs"

// @Author: Feng
// @Date: 2022/3/25 17:18

const (
	errorParam = 10000 + iota
	errorHandler
)

var (
	ParamError   = errs.NewError(errorParam, "参数错误")
	HandlerError = errs.NewError(errorHandler, "该appid尚未注册")
)
