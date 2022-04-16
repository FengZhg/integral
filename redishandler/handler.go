package redishandler

import (
	"integral/logic"
)

// @Author: Feng
// @Date: 2022/3/25 17:46

type RedisHandler struct{}

func init() {
	// 注册基于redis的积分处理器
	logic.RegisterIntegralHandler("10000", &RedisHandler{})
}

// lua 脚本传入参数统一格式
// KEYS[1]:余额
// KEYS[2]:订单标志key
// KEYS[3]:回滚标志
