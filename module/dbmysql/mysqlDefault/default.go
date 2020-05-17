package mysqlDefault

import (
	"github.com/bridgewwater/golang-project-temple-docker-db/module/optgormmysql"
	"github.com/bridgewwater/golang-project-temple-docker-db/pkg/zlog"
)

const name string = "default"

var defaultOptMysql *optgormmysql.OptMysql

func Init() error {
	if defaultOptMysql == nil {
		config := optgormmysql.NewConfig(
			optgormmysql.WithName(name),
		)
		optMysql := optgormmysql.NewMySQL(*config).InitByName()
		zlog.S().Infof("DB mysql mark: %v, use conn: %v", optMysql.Name(), optMysql.ConnectionURI())
		if optMysql.ErrClientConnection != nil {
			return optMysql.ErrClientConnection
		}

		defaultOptMysql = &optMysql
	}
	return nil
}

func Opt() *optgormmysql.OptMysql {
	return defaultOptMysql
}
