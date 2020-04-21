package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

const (
	// env prefix is web
	defaultEnvPrefix string = "ENV_WEB"
	// env ENV_WEB_HTTPS_ENABLE default false
	defaultEnvHttpsEnable string = "HTTPS_ENABLE"
	// env ENV_WEB_HOST default ""
	defaultEnvHost string = "HOST"
	// env ENV_AUTO_HOST default true
	defaultEnvAutoGetHost string = "AUTO_HOST"
)

var mustConfigString = []string{
	"run_mode",
	"addr",
	"name",
	"api_version",
	// project set
}

// this config can config by ENV, and set writeLogFile true
//	ENV_WEB_HTTPS_ENABLE=false
//	ENV_AUTO_HOST=true
//	ENV_WEB_HOST 127.0.0.1:8000
// config set by
//	cfgPath path for config, read default config by conf/config.yaml
func Init(cfgPath string) error {
	return InitCfg(cfgPath, true)
}

// this config can config by ENV
//	ENV_WEB_HTTPS_ENABLE=false
//	ENV_AUTO_HOST=true
//	ENV_WEB_HOST 127.0.0.1:8000
// config set by
//	cfgPath path for config, read default config by conf/config.yaml
//	writeLogFile true will write log
func InitCfg(cfgPath string, writeLogFile bool) error {
	c := ConfFile{
		ConfigPath:   cfgPath,
		WriteLogFile: writeLogFile,
	}

	// initialize configuration file
	if err := c.initConfig(); err != nil {
		return err
	}

	// initialization log package
	if err := c.initLog(); err != nil {
		return err
	}

	// init baseConfig
	initBaseConf()

	// TODO other config

	// monitor configuration changes and hot loaders
	c.watchConfig()

	return nil
}

func (c *ConfFile) initConfig() error {
	if c.ConfigPath != "" {
		viper.SetConfigFile(c.ConfigPath) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath(filepath.Join("conf")) // 如果没有指定配置文件，则解析默认的配置文件 conf/config.go
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")          // 设置配置文件格式为YAML
	viper.AutomaticEnv()                 // 读取匹配的环境变量
	viper.SetEnvPrefix(defaultEnvPrefix) // 读取环境变量的前缀为 defaultEnvPrefix

	// 设置默认环境变量
	_ = os.Setenv(defaultEnvHost, "")
	_ = os.Setenv(defaultEnvHttpsEnable, "false")
	_ = os.Setenv(defaultEnvAutoGetHost, "true")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	if err := checkMustHasString(); err != nil {
		return err
	}

	err := viper.Unmarshal(&global)
	if err != nil {
		return err
	}

	return nil
}

// check config.yaml must has string key
//	config.mustConfigString
func checkMustHasString() error {
	for _, config := range mustConfigString {
		if "" == viper.GetString(config) {
			return fmt.Errorf("not has must string key [ %v ]", config)
		}
	}
	return nil
}
