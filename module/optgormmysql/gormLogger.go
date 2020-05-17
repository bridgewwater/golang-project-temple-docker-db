package optgormmysql

import (
	"github.com/go-sql-driver/mysql"
	"reflect"
	"time"
)

type GormLoggerHandler interface {
	Default(values ...interface{})
	SQLPrint(source interface{}, timeCost time.Duration, sql string)
	LogPrint(source interface{}, typeMark string)
	MySQLErrorPrint(source interface{}, error *mysql.MySQLError)
}

type GormLogger struct {
	IsOpen  bool
	Handler GormLoggerHandler
}

func (logger *GormLogger) Print(values ...interface{}) {
	if !logger.IsOpen {
		return
	}

	if logger.Handler != nil {
		var (
			level    = values[0]
			source   = values[1]
			typeMark = reflect.TypeOf(values[2])
		)
		switch level {
		default:
			logger.Handler.Default(values)
		case "sql":
			sql := values[3].(string)
			timeCost := values[2].(time.Duration)
			logger.Handler.SQLPrint(source, timeCost, sql)
		case "log":
			switch typeMark.String() {
			default:
				logger.Handler.LogPrint(source, typeMark.String())
			case "*mysql.MySQLError":
				sqlError := values[2].(*mysql.MySQLError)
				logger.Handler.MySQLErrorPrint(source, sqlError)
			}
		}
	}
}
