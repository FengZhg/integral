package dbhandler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"integral/dao"
	"integral/model"
	"integral/server"
	"integral/utils"
	"time"
)

// @Author: Feng
// @Date: 2022/4/7 14:59

func (d *dbHandler) Modify(ctx *gin.Context, req *server.ModifyReq, rsp *server.ModifyRsp) error {
	// 修改积分
	err := dao.ExecTransaction(ctx, getModifyTransaction(ctx, req))
	if err != nil {
		log.Errorf("Exec Modify Tx Error %v", err)
		return err
	}
	return nil
}

//getModifyTransaction 获取执行事务中的具体操作的闭包
func getModifyTransaction(ctx *gin.Context, req *server.ModifyReq) func(*gin.Context, *sql.Tx) error {
	// 处理预差值
	difBalance := req.GetIntegral()
	if req.GetOpt() == model.DescType {
		difBalance = -1 * difBalance
	}
	now := time.Now()
	// 构造闭包
	return func(context *gin.Context, tx *sql.Tx) error {
		// 修改余额
		modifySql := fmt.Sprintf("inset into DBIntegral_%v.tbIntegral_%v(appid,type,id,integral,desc) "+
			"value(?,?,?,?,?) on duplicate key update set integral = integral + ? where appid = ? and type = ? and id"+
			" = ?;",
			req.GetAppid(), utils.GetDBIndex(req.GetUid()))
		_, err := tx.Exec(modifySql, req.GetAppid(), req.GetType(), req.GetUid(), difBalance, req.GetDesc(),
			difBalance, req.GetAppid(), req.GetType(), req.GetUid())
		if err != nil {
			return err
		}

		// 插入流水
		insertFlowSql := fmt.Sprintf("insert into DBIntegralFlow_%v.tbIntegralFlow_%v(id,"+
			"oid,appid,type,opt,integral,timestamp,time,desc) value(?,?,?,?,?,?,?,?,?);", req.GetAppid(),
			utils.GetDBIndex(req.GetUid()))
		_, err = tx.Exec(insertFlowSql, req.GetUid(), req.GetOid(), req.GetAppid(), req.GetType(), req.GetOpt(), req.GetIntegral(),
			now.UnixNano(), now.Format("2006-01-02 15:04:05"), req.GetDesc())
		if err != nil {
			return err
		}
		return nil
	}
}
