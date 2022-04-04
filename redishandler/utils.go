package redishandler

import "fmt"

// @Author: Feng
// @Date: 2022/4/4 17:00
// 只有redis handler 用到的工具函数
// tid即type

//getOrderKey 获取订单键
func getOrderKey(appid, tid, uid, oid string) string {
	return fmt.Sprintf("oid_%v_%v_%v_{%v}", appid, tid, oid, uid)
}

//getFlowKey 获取流水键
func getFlowKey(appid, tid, uid, oid string) string {
	return fmt.Sprintf("flow_%v_%v_%v_{%v}", appid, tid, oid, uid)
}

//getBalanceKey 获取订单键
func getBalanceKey(appid, tid, uid string) string {
	return fmt.Sprintf("balance_%v_%v_{%v}", appid, tid, uid)
}
