package pulsarClient

import (
	"integral/dao"
	"integral/model"
)

// @Author: Feng
// @Date: 2022/4/22 12:54

var (
	//PulsarOpt 参数情况
	PulsarOpt = NewPulsarOptions(
		WithUrl(model.PulsarUrl),
		WithTopic(model.PulsarTopic),
	)
)

// 初始化
func init() {
	// 启动消费者守护协程
	go NewPulsarConsumerDaemon(
		PulsarOpt,
		dao.FlowConsumeCallback,
	).Start()
}
