package dbhandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"integral/dao"
	"integral/server"
	"integral/utils"
)

// @Author: Feng
// @Date: 2022/4/11 20:20

//Query 余额修改
func (d *dbHandler) Query(ctx *gin.Context, req *server.QueryReq, rsp *server.QueryRsp) error {
	for _, uid := range req.GetUids() {
		subRsp, err := doQuery(ctx, uid, req.GetType(), req.GetAppid())
		if err != nil {
			log.Errorf("Do Query User Error %v", err)
			continue
		}
		rsp.UsersRsp = append(rsp.UsersRsp, subRsp)
	}
	return nil
}

//doQuery 进行查询
func doQuery(ctx *gin.Context, uid, tid, appid string) (*server.SingleQueryRsp, error) {
	//构造查询语句
	querySql := fmt.Sprintf("select integral from DBIntegral_%v.tbIntegral_%v where appid = ? and type = ? id = ?;",
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
	return &server.SingleQueryRsp{
		Appid:    appid,
		Type:     tid,
		Integral: integral,
		Uid:      uid,
	}, nil
}
