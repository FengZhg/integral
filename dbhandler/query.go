package dbhandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"integral/dao"
	"integral/model"
	"integral/utils"
	"sync"
)

// @Author: Feng
// @Date: 2022/4/11 20:20

//Query 余额修改
func (d *dbHandler) Query(ctx *gin.Context, req *model.QueryReq, rsp *model.QueryRsp) error {
	wg := sync.WaitGroup{}
	queryRsps := make([]*model.SingleQueryRsp, len(req.GetUids()))
	for index, uid := range req.GetUids() {
		wg.Add(1)
		go func(idx int, id string) {
			defer wg.Done()
			subRsp, err := doQuery(ctx, uid, req.GetType(), req.GetAppid())
			if err != nil {
				log.Errorf("Query Single User Balance Error %v", err)
				return
			}
			queryRsps[idx] = subRsp
		}(index, uid)
	}
	wg.Wait()
	rsp.UsersRsp = queryRsps
	return nil
}

//doQuery 进行查询
func doQuery(ctx *gin.Context, uid, tid, appid string) (*model.SingleQueryRsp, error) {
	//构造查询语句
	querySql := fmt.Sprintf("select integral from DBIntegral_%v.tbIntegral_%v where appid = ? and type = ? and id = ?;",
		appid, utils.GetDBIndex(uid))
	row := dao.GetDBClient().QueryRow(querySql, appid, tid, uid)
	if row.Err() != nil {
		return nil, row.Err()
	}
	// 获取结果
	var integral int64
	err := row.Scan(&integral)
	if err != nil {
		return nil, err
	}
	return &model.SingleQueryRsp{
		Appid:    appid,
		Type:     tid,
		Integral: integral,
		Uid:      uid,
	}, nil
}
