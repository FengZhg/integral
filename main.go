package main

import (
	_ "github.com/FengZhg/go_tools/gin_logrus"
	"integral/dao"
	"integral/dao/pulsarClient"
	"integral/model"
	"integral/server"
)

// @Author: Feng
// @Date: 2022/3/24 17:50

func main() {
	engine := server.NewServer()
	engine.Run("0.0.0.0:10000")
}

// 初始化
func init() {
	initPulsarConsumer()
}

func initPulsarConsumer() {
	// 启动消费者守护协程
	pulsarClient.NewPulsarConsumerDaemon(
		model.PulsarOpt,
		dao.FlowConsumeCallback,
	).Start()
}
