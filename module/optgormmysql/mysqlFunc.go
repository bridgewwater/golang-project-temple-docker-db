package optgormmysql

import (
	"fmt"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/optgormmysql/optcfgmysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type OptMysql struct {
	errClientNotInit error

	name            string
	mysqlConnection string

	gormClient *gorm.DB

	openLog             bool
	ErrClientConnection error
}

func (o OptMysql) InitByName() OptMysql {
	if mysqlConfigList == nil {
		panic(errMysqlListListEmpty)
	}
	mySQLCfg := optcfgmysql.ByName(*mysqlConfigList, o.name)

	openLog := parseEnvBooleanOrFalse(parseEnvByName(o.name, envKeyMysqlLog))
	dbHost := parseEnvStringOrDefault(parseEnvByName(o.name, envKeyMysqlHost), mySQLCfg.Host)
	dbName := parseEnvStringOrDefault(parseEnvByName(o.name, envKeyMysqlDBName), mySQLCfg.DBName)
	dbUser := parseEnvStringOrDefault(parseEnvByName(o.name, envKeyMysqlDBUser), mySQLCfg.Username)
	dbPassword := parseEnvStringOrDefault(parseEnvByName(o.name, envKeyMysqlDBPWD), mySQLCfg.Password)
	dbConfig := parseEnvStringOrDefault(parseEnvByName(o.name, envKeyMysqlConfig), mySQLCfg.Config)
	dbMaxIdConns := parseEnvIntOrDefault(parseEnvByName(o.name, envKeyMysqlMaxIdleConns), mySQLCfg.MaxIdleConns)
	dbMaxOpenConns := parseEnvIntOrDefault(parseEnvByName(o.name, envKeyMysqlMaxOpenConns), mySQLCfg.MaxOpenConns)
	dbConnMaxLifetime := parseEnvIntOrDefault(parseEnvByName(o.name, envKeyMysqlConnMaxLifetime), mySQLCfg.ConnMaxLifetime)

	mysqlConnection := fmt.Sprintf("%v:%v@(%v)/%v?%v",
		dbUser, dbPassword, dbHost, dbName, dbConfig)
	o.mysqlConnection = mysqlConnection

	db, err := gorm.Open("mysql", o.mysqlConnection)
	if err != nil {
		o.ErrClientConnection = err
		return o
	}
	db.DB().SetMaxIdleConns(dbMaxIdConns)
	db.DB().SetMaxOpenConns(dbMaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(dbConnMaxLifetime) * time.Minute)
	o.gormClient = db
	o.openLog = openLog
	return o
}

func (o OptMysql) Logger(logger GormLogger) error {
	if o.openLog {
		o.gormClient.SetLogger(&logger)
		return nil
	}
	return fmt.Errorf("not set openlog, can use env: [ %v ] or by config: open_log", parseEnvByName(o.name, envKeyMysqlLog))
}

func (o OptMysql) Name() string {
	return o.name
}

func (o OptMysql) ConnectionURI() string {
	return o.mysqlConnection
}

func (o OptMysql) Client() *gorm.DB {
	return o.gormClient
}

type MysqlFunc interface {
	InitByName() OptMysql
	Logger(GormLogger) error
	Name() string
	ConnectionURI() string
	Client() *gorm.DB
}

func NewMySQL(config Config) MysqlFunc {
	return &OptMysql{
		name: config.Name,

		errClientNotInit: fmt.Errorf("%v opt mysql : client empty, must be init", config.Name),
	}
}
