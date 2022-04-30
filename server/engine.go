package server

import (
	"github.com/FengZhg/go_tools/gin_middleware"
	"github.com/gin-gonic/gin"
	"integral/logic"
	"net/http"
	"os/exec"
)

// @Author: Feng
// @Date: 2022/3/24 17:53

func init() {
	exec.Command("")
}

func NewServer() *gin.Engine {
	// 生成默认
	engine := gin.Default()

	// 中间件
	// 超时控制中间件
	engine.Use(gin_middleware.NewRequestLog(nil).RequestLogMiddleware())
	engine.Use(gin_middleware.ReplyMiddleware())
	engine.Use(gin_middleware.TimeoutMiddleware())

	// 校验登录态的接口
	api := engine.Group("api").Use(logic.Jwt.AuthMiddleware())
	{
		api.POST("/query", queryBase)
		api.POST("/modify", modifyBase)
		api.POST("/rollback", rollbackBase)
		api.POST("/queryflow", queryFlowBase)
	}

	engine.POST("/token", generateTokenBase)

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Server Started")
	})
	return engine
}
