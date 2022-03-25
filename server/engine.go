package server

import (
	"github.com/FengZhg/go_tools/gin_middleware"
	"github.com/FengZhg/go_tools/goJwt"
	"github.com/gin-gonic/gin"
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
	engine.Use(gin_middleware.NewRequestLog(nil, nil).RequestLogMiddleware())
	engine.Use(gin_middleware.ReplyMiddleware())
	engine.Use(gin_middleware.TimeoutMiddleware())

	// 校验登录态的接口
	api := engine.Group("api").Use(goJwt.NewES512().AuthMiddleware())
	{
		api.POST("Query", queryBase)
		api.POST("Modify", modifyBase)
		api.POST("Rollback", rollbackBase)
		api.POST("QueryFlow", queryFlowBase)
	}

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Server Started")
	})
	return engine
}
