package redishandler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"integral/dao"
	"integral/model"
)

// @Author: Feng
// @Date: 2022/3/26 14:08

//Query Redis处理器的积分查询函数
func (r *RedisHandler) Query(ctx *gin.Context, req *model.QueryReq, rsp *model.QueryRsp) error {
	var queryRsps []*model.SingleQueryRsp
	for _, uid := range req.GetUids() {
		balance, err := singleQuery(ctx, req.GetAppid(), req.GetType(), uid)
		if err != nil {
			log.Errorf("Query Single User Balance Error %v", err)
			continue
		}
		queryRsps = append(queryRsps, &model.SingleQueryRsp{
			Uid:      uid,
			Appid:    req.GetAppid(),
			Type:     req.GetType(),
			Integral: balance,
		})
	}
	rsp.UsersRsp = queryRsps
	return nil
}

//singleQuery 请求一个用户的余额
func singleQuery(ctx *gin.Context, appid, tid, uid string) (int64, error) {
	balance, err := dao.GetRedisClient().Get(ctx, getBalanceKey(appid, tid, uid)).Int64()
	if err != nil {
		log.Errorf("Query Single User Error %v", err)
		return 0, err
	}
	return balance, nil
}
