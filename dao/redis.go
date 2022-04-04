package dao

import (
	"github.com/go-redis/redis/v8"
	"integral/model"
)

// @Author: Feng
// @Date: 2022/4/4 16:25

var redisCli *redis.Client = nil

//GetRedisClient 构造redis客户端
func GetRedisClient() *redis.Client {
	if redisCli == nil {
		redisCli = redis.NewClient(&redis.Options{
			Addr:     model.RedisAddr,
			Password: model.RedisPassword,
			DB:       model.RedisDB,
		})
	}
	return redisCli
}
