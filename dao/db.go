package dao

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"integral/logic"
	"integral/model"
	"integral/utils"
	"time"
)

// @Author: Feng
// @Date: 2022/3/28 15:47

var dbCli *sql.DB = nil

func init() {
	var err error
	dbCli, err = newDBClient()
	if err != nil {
		log.Errorf("Init Db Client Error %v", err)
		panic(err)
	}
}

//newDBClient 创建全局链接
func newDBClient() (*sql.DB, error) {
	//创建数据库连接
	db, err := sql.Open(model.DBDriverName, model.DBDriverConfigStr)
	if err != nil {
		log.Errorf("MySQL Open Conn Error err = %v", err)
		return nil, err
	}
	return db, nil
}

//GetDBClient 获取数据库链接
func GetDBClient() *sql.DB {
	return dbCli
}

//FlowConsumeCallback 用于消费pulsar的
func FlowConsumeCallback(msg pulsar.ConsumerMessage) error {
	flow := &logic.SingleFlow{}
	// 反序列化语句
	err := json.Unmarshal(msg.Payload(), flow)
	if err != nil {
		log.Errorf("Parse Pulsar Message Error %v", err)
		return err
	}

	// 构造请求语句
	insertFlowSql := fmt.Sprintf("insert into DBIntegralFlow_%v.tbIntegralFlow_%v(id,"+
		"oid,appid,type,opt,integral,timestamp,time) value(?,?,?,?,?,?,?,?,?);", flow.GetAppid(), utils.GetDBIndex(flow.GetUid()))
	param := []interface{}{
		flow.GetUid(), flow.GetOid(), flow.GetAppid(), flow.GetType(), flow.GetOpt(), flow.GetIntegral(), flow.GetTimestamp(), flow.GetTime(),
	}

	// 生成带超时的ctx
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = dbCli.ExecContext(ctx, insertFlowSql, param...)
	if err != nil {
		log.Errorf("Mysql Insert Error %v", err)
		return err
	}
	return nil
}

//ExecTransaction 执行mysql语句
func ExecTransaction(ctx *gin.Context, handler func(*gin.Context, *sql.Tx) error) error {
	// 开始事务
	tx, err := dbCli.Begin()
	if err != nil {
		log.Errorf("Transaction Begin Error err = %v", err)
		return err
	}

	err = handler(ctx, tx)
	if err != nil {
		log.Errorf("Run Transaction Error err = %v", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
