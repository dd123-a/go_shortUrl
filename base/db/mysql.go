package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"short_url/base/config"
	"short_url/base/tool"
)

func initMysql() {
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8",
		config.GetMysqlConfig().GetName(), config.GetMysqlConfig().GetPass(), config.GetMysqlConfig().GetIP(), config.GetMysqlConfig().GetPort(), config.GetMysqlConfig().GetDb())
	//初始化mysql引擎
	if mysqlDb, err = xorm.NewEngine("mysql", url); err != nil {
		panic(err)
	}
	mysqlDb.SetMaxOpenConns(config.GetMysqlConfig().GetMaxIdle())
	mysqlDb.SetMaxIdleConns(config.GetMysqlConfig().GetMaxOpen())
	_ =mysqlDb.Ping()
	if config.GetToolLogConfig().GetDevelopment() {
		mysqlDb.ShowSQL(true)
		mysqlDb.ShowExecTime(true)
	}
	tool.GetLogger().Debug("mysql : " + url)
}


func closeMysql() {
	if mysqlDb != nil {
		_ = mysqlDb.Close()
	}
}