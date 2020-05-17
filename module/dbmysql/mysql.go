package dbmysql

import (
	"github.com/bridgewwater/golang-project-temple-docker-db/cfg"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/dbmysql/mysqlDefault"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/optgormmysql"
)

func InitMySQLOpt() error {
	err := optgormmysql.InitByConfigList(cfg.Global().MySQLConfigList)
	if err != nil {
		return err
	}
	return mysqlDefault.Init()
}
