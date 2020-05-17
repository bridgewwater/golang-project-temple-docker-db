package optgormmysql

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

const (
	envKeyMysqlPrefix          = "DB_MYSQL"
	envKeyMysqlLog             = "OPEN_LOG"
	envKeyMysqlHost            = "HOST"
	envKeyMysqlDBName          = "DBNAME"
	envKeyMysqlDBUser          = "USER"
	envKeyMysqlDBPWD           = "PWD"
	envKeyMysqlConfig          = "CONFIG"
	envKeyMysqlMaxIdleConns    = "MAX_IDLE_CONNS"
	envKeyMysqlMaxOpenConns    = "MAX_OPEN_CONNS"
	envKeyMysqlConnMaxLifetime = "CONN_MAX_LIFETIME"
)

type Config struct {
	Name string
}

var defaultConfig = setDefaultConfig()

type ConfigOption func(*Config)

func setDefaultConfig() *Config {
	return &Config{
		Name: "default",
	}
}

func NewConfig(opts ...ConfigOption) (opt *Config) {
	opt = defaultConfig
	for _, o := range opts {
		o(opt)
	}
	defaultConfig = setDefaultConfig()
	return
}

func WithName(name string) ConfigOption {
	return func(o *Config) {
		o.Name = name
	}
}

func NewConfigFull() *Config {
	return setDefaultConfig()
}

func parseEnvBooleanOrFalse(envKey string) bool {
	return viper.GetBool(envKey)
}

func parseEnvStringOrDefault(envKey, defaultStr string) string {
	var result string
	if viper.GetString(envKey) == "" {
		result = defaultStr
	} else {
		result = viper.GetString(envKey)
	}
	return result
}

func parseEnvIntOrDefault(envKey string, defaultInt int) int {
	var result int
	if viper.GetInt(envKey) == 0 {
		result = defaultInt
	} else {
		result = viper.GetInt(envKey)
	}
	return result
}

func parseEnvByName(name string, last string) string {
	return fmt.Sprintf("%s_%s_%s", envKeyMysqlPrefix, strings.ToUpper(name), last)
}
