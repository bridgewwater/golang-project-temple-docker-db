package cfg

import (
	"github.com/bridgewwater/golang-project-temple-docker-db/pkg/zlog"
	"github.com/spf13/viper"
)

var baseConf *baseConfig

type baseConfig struct {
	IsDebug    bool
	EnvName    string
	ApiVersion string
	SSLEnable  bool
}

func BaseConfig() *baseConfig {
	return baseConf
}

func EnvName() string {
	return baseConf.EnvName
}

func IsDebug() bool {
	return baseConf.IsDebug
}

// read default config by conf/config.yaml
// can change by CLI by `-c`
// this config can config by ENV
//	ENV_WEB_HTTPS_ENABLE=false
//	ENV_AUTO_HOST=true
//	ENV_WEB_HOST 127.0.0.1:8000
func initBaseConf() {
	var env = viper.GetString("run_mode")
	var confDebug bool
	if env == "debug" || env == "test" {
		confDebug = true
	} else {
		confDebug = false
	}

	ssLEnable := false
	if viper.GetBool(defaultEnvHttpsEnable) {
		ssLEnable = true
	} else {
		ssLEnable = viper.GetBool("sslEnable")
	}

	zlog.S().Debugf("config.BaseConfig()\nIsDebug: %v\nEnvName: %v\nApiVersion: %v\nSSLEnable: %v",
		confDebug, env, viper.GetString("api_version"), ssLEnable)
	baseConf = &baseConfig{
		IsDebug:    confDebug,
		EnvName:    env,
		ApiVersion: viper.GetString("api_version"),
		SSLEnable:  ssLEnable,
	}
}
