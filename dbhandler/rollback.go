package dbhandler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"integral/model"
	"integral/server"
	"integral/utils"
)

// @Author: Feng
// @Date: 2022/4/11 20:21

//Rollback Redis处理器回滚
func (d *dbHandler) Rollback(ctx *gin.Context, req *server.RollbackReq, rsp *server.RollbackRsp) error {

	return nil
}

//getRollbackTransaction
func getRollbackTransaction(ctx *gin.Context, req *server.RollbackReq) func(*gin.Context, *sql.Tx) error {
	// 预构造flow结构体
	flow := &server.SingleFlow{}

	return func(ctx *gin.Context, tx *sql.Tx) error {
		// 查询回滚记录
		queryFlowSql := fmt.Sprintf("select opt,integral,timestamp,rollback from DBIntegralFlow_%v.tbIntegralFlow_%v where appid = ? and type = ? and id = ? and oid = ?;", req.GetAppid(), utils.GetDBIndex(req.GetUid()))
		row := tx.QueryRowContext(ctx, queryFlowSql, req.GetAppid(), req.GetType(), req.GetUid(), req.GetOid())
		if row.Err() != nil {
			return row.Err()
		}
		// 读入
		err := row.Scan(&flow.Opt, &flow.Integral, &flow.Timestamp, &flow.Rollback)
		if err != nil {
			return err
		}
		// 判断是否已经回滚
		if flow.GetRollback() == true {
			return model.AlreadyRollbackError
		}
		// 处理回滚需要的数据
		integral := flow.GetIntegral()
		if flow.GetOpt() == 1 {
			integral = -1 * integral
		}
		rollback
		return nil
	}
}
