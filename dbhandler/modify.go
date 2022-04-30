package dbhandler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"integral/dao"
	"integral/model"
	"integral/utils"
	"time"
)

// @Author: Feng
// @Date: 2022/4/7 14:59

func (d *dbHandler) Modify(ctx *gin.Context, req *model.ModifyReq, rsp *model.ModifyRsp) error {
	// 处理预差值
	difBalance := req.GetIntegral()
	if req.GetOpt() == model.DescType {
		difBalance = -1 * difBalance
	}
	now := time.Now()
	// 构造闭包
	txCallback := func(context *gin.Context, tx *sql.Tx) error {
		// 修改余额
		modifySql := fmt.Sprintf("insert into DBIntegral_%v.tbIntegral_%v(appid,type,id,integral) "+
			"value(?,?,?,?) on duplicate key update integral = integral + ?;", req.GetAppid(), utils.GetDBIndex(req.GetUid()))
		_, err := tx.Exec(modifySql, req.GetAppid(), req.GetType(), req.GetUid(), difBalance, difBalance)
		if err != nil {
			return err
		}
		// 插入流水
		insertFlowSql := fmt.Sprintf("insert into DBIntegralFlow_%v.tbIntegralFlow_%v(id,"+
			"oid,appid,type,opt,integral,timestamp,time) value(?,?,?,?,?,?,?,?);", req.GetAppid(),
			utils.GetDBIndex(req.GetUid()))
		_, err = tx.Exec(insertFlowSql, req.GetUid(), req.GetOid(), req.GetAppid(), req.GetType(), req.GetOpt(), req.GetIntegral(),
			now.UnixNano(), now.Format("2006-01-02 15:04:05"))
		//if err != nil {
		//	return err
		//}
		// 查询返回余额
		querySql := fmt.Sprintf("select integral from DBIntegral_%v.tbIntegral_%v where appid = ? and type = ? "+
			"and id = ?", req.GetAppid(), utils.GetDBIndex(req.GetUid()))
		row := tx.QueryRow(querySql, req.GetAppid(), req.GetType(), req.GetUid())
		if row.Err() != nil {
			return row.Err()
		}
		return row.Scan(&rsp.Integral)
	}
	// 修改积分
	err := dao.ExecTransaction(ctx, txCallback)
	if err != nil {
		log.Errorf("Exec Modify Tx Error %v", err)
		return err
	}
	return nil
}
