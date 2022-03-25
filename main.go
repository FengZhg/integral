package main

import (
	_ "github.com/FengZhg/go_tools/gin_logrus"
	"integral/server"
)

// @Author: Feng
// @Date: 2022/3/24 17:50

func main() {
	engine := server.NewServer()
	engine.Run("0.0.0.0:10000")
}
