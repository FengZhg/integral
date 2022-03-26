package redishandler

import "integral/logic"

// @Author: Feng
// @Date: 2022/3/25 17:46

type redisHandler struct{}

func init() {
	// 注册基于redis的积分处理器
	logic.RegisterIntegralHandler("10000", &redisHandler{})
}
