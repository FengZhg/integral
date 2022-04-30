package dbhandler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"integral/dao"
	"integral/model"
	"integral/utils"
)

// @Author: Feng
// @Date: 2022/4/11 20:21

//Rollback Redis处理器回滚
func (d *dbHandler) Rollback(ctx *gin.Context, req *model.RollbackReq, rsp *model.RollbackRsp) error {
	// 执行回滚事务
	err := dao.ExecTransaction(ctx, getRollbackTransaction(ctx, req))
	if err != nil {
		log.Errorf("Exec Transaction Error %v", err)
		return err
	}
	return nil
}

//getRollbackTransaction
func getRollbackTransaction(ctx *gin.Context, req *model.RollbackReq) func(*gin.Context, *sql.Tx) error {
	// 预构造flow结构体
	flow := &model.SingleFlow{}

	return func(ctx *gin.Context, tx *sql.Tx) error {
		// 查询回滚记录
		queryFlowSql := fmt.Sprintf("select opt,integral,rollback from DBIntegralFlow_%v.tbIntegralFlow_%v "+
			"where appid = ? and type = ? and id = ? and oid = ? for update;", req.GetAppid(),
			utils.GetDBIndex(req.GetUid()))
		row := tx.QueryRowContext(ctx, queryFlowSql, req.GetAppid(), req.GetType(), req.GetUid(), req.GetOid())
		if row.Err() != nil {
			return row.Err()
		}
		// 读入
		err := row.Scan(&flow.Opt, &flow.Integral, &flow.Rollback)
		if err != nil {
			return err
		}
		// 判断是否已经回滚
		if flow.GetRollback() == true {
			return model.AlreadyRollbackError
		}
		// 修改流水标志
		updateFlowSql := fmt.Sprintf("update DBIntegralFlow_%v.tbIntegralFlow_%v set rollback = true "+
			"where appid = ? and type = ? and id = ? and oid = ? and rollback = false;", req.GetAppid(), utils.GetDBIndex(req.GetUid()))
		re, err := tx.ExecContext(ctx, updateFlowSql, req.GetAppid(), req.GetType(), req.GetUid(), req.GetOid())
		if row.Err() != nil {
			return row.Err()
		}
		aff, err := re.RowsAffected()
		if err != nil {
			return err
		}
		if aff == 0 {
			return model.UpdateUnexpectedError
		}
		// 处理回滚需要的数据
		integral := int64(flow.GetIntegral())
		if flow.GetOpt() == 1 {
			integral = -1 * integral
		}
		// 更新余额
		updateIntegralSql := fmt.Sprintf("update DBIntegral_%v.tbIntegral_%v set integral = integral + ? "+
			"where appid = ? and type = ? and id = ?", req.GetAppid(), utils.GetDBIndex(req.GetUid()))
		re, err = tx.ExecContext(ctx, updateIntegralSql, integral, req.GetAppid(), req.GetType(), req.GetUid())
		if err != nil {
			return err
		}
		aff, err = re.RowsAffected()
		if err != nil {
			return err
		}
		if aff == 0 {
			return model.UpdateUnexpectedError
		}
		return nil
	}
}
