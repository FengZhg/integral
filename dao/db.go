package dao

import (
	"database/sql"
	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
	"integral/model"
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

//flowConsumeCallback 用于消费pulsar的
func flowConsumeCallback(msg *pulsar.Message) {

}
