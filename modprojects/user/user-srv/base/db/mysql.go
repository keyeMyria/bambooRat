package db

import (
	"database/sql"
	"github.com/micro/go-log"
	"github.com/xuyiwenak/bambooRat/modprojects/user/user-srv/base/config"
)

func initMysql() {

	var err error

	log.Log("sql--->", config.GetMysqlConfig().GetURL())

	// 创建连接
	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Log("sql--->ok")
}
