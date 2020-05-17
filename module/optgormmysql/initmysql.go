package optgormmysql

import (
	"fmt"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/optgormmysql/optcfgmysql"
)

var mysqlConfigList *[]optcfgmysql.Config

var errMysqlListListEmpty = fmt.Errorf("mysql config list is empty, you must use optgormmysql.InitByConfigList()")

// init MySQL by yaml list for config
//	mysql_clients:
//  - name: default
//    open_log: true
//    host: localhost:3306
//    db_name: default
//    username: user
//    password: password
//    config: charset=utf8&parseTime=True&loc=UTC
//    max_idle_conns: 10
//    max_open_conns: 100
//    conn_max_lifetime: 0
func InitByConfigList(list []optcfgmysql.Config) error {
	if list == nil {
		return errMysqlListListEmpty
	}
	if len(list) == 0 {
		return errMysqlListListEmpty
	}
	mysqlConfigList = &list
	return nil
}
