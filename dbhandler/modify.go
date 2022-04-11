package dbhandler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"integral/dao"
	"integral/server"
)

// @Author: Feng
// @Date: 2022/4/7 14:59

func (d *dbHandler) Modify(ctx *gin.Context, req *server.ModifyReq, rsp *server.ModifyRsp) error {

	return nil
}

//doModify 修改战功
func doModify(ctx *gin.Context, req *server.ModifyReq) error {
	// 获取连接启动事务
	dbCli := dao.GetDBClient()
	tx, err := dbCli.Begin()
	if err != nil {
		log.Errorf("Begin Transaction Error %v", err)
		return err
	}

	err = doModifyTx(ctx, tx)
	if err != nil {
		log.Errorf("Exec Modify Tx Error %v", err)
		tx.Rollback()
		return err
	}

	// 提交事务
	tx.Commit()
	return nil
}

//doModifyTx 执行事务中的具体操作
func doModifyTx(ctx *gin.Context, tx *sql.Tx) error {
	// 修改余额
	tx.Exec()

	// 插入流水

	return nil
}
