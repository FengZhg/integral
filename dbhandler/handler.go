package dbhandler

import "integral/logic"

// @Author: Feng
// @Date: 2022/4/7 14:58

type dbHandler struct{}

func init() {
	// 注册基于db的积分处理器
	logic.RegisterIntegralHandler("10001", &dbHandler{})
}
