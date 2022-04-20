package logic

import (
	"github.com/FengZhg/go_tools/goJwt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/4/19 15:52

var Jwt = goJwt.NewES512()

//GenerateToken 分配
func GenerateToken(ctx *gin.Context, rsp map[string]string) error {
	token, err := Jwt.ApplyToken("lcx", "lcx")
	if err != nil {
		log.Errorf("Apply Token Error %v", err)
		return err
	}
	rsp["token"] = token
	return nil
}
