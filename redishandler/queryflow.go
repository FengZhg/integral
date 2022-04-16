package redishandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"integral/dao"
	"integral/server"
	"integral/utils"
)

// @Author: Feng
// @Date: 2022/3/26 14:08

//QueryFlow Redis处理器查询积分流水
func (r *RedisHandler) QueryFlow(ctx *gin.Context, req *server.QueryFlowReq, rsp *server.QueryFlowRsp) error {

	// 查询流水
	flows, err := queryFlow(ctx, req)
	if err != nil {
		log.Errorf("Query Flow Error %v", err)
		return err
	}
	rsp.Flows = flows
	return nil
}

//queryFlow 查询积分流水
func queryFlow(ctx *gin.Context, req *server.QueryFlowReq) ([]*server.SingleFlow, error) {
	// 构造
	var flows []*server.SingleFlow
	query := fmt.Sprintf("select id,oid,appid,type,opt,integral,timestamp,time,"+
		"desc,rollback from DBIntegralFlow_%v.tbIntegralFlow_%v where id=? order by timestamp desc limit ?,?",
		req.GetAppid(), utils.GetDBIndex(req.GetUid()))
	param := []interface{}{req.GetUid(), req.GetOffset(), req.GetNum()}

	// 请求
	rows, err := dao.GetDBClient().QueryContext(ctx, query, param...)
	if err != nil {
		log.Errorf("Query Flow Error %v", err)
		return nil, err
	}

	// 解析返回
	for rows.Next() {
		f := &server.SingleFlow{}
		err := rows.Scan(&f.Uid, &f.Oid, &f.Appid, &f.Type, f.Opt, &f.Integral, &f.Timestamp, &f.Time, &f.Desc, &f.Rollback)
		if err != nil {
			continue
		}
		flows = append(flows, f)
	}
	return flows, nil
}
