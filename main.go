package main

import (
	_ "github.com/FengZhg/go_tools/gin_logrus"
	log "github.com/sirupsen/logrus"
	_ "integral/dao/pulsarClient"
	_ "integral/dbhandler"
	_ "integral/redishandler"
	"integral/server"
)

// @Author: Feng
// @Date: 2022/3/24 17:50

func main() {
	log.SetLevel(log.ErrorLevel)
	engine := server.NewServer()
	err := engine.Run("0.0.0.0:10000")
	if err != nil {
		panic(err)
	}
}
