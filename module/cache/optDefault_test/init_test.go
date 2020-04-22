package optDefault_test

import (
	"fmt"
	"github.com/bridgewwater/golang-project-temple-docker-db/cfg"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/cache"
	"github.com/sinlovgo/timezone"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// setup
	initOptRedisForTest()

	os.Exit(m.Run())
	// teardown
}

func initOptRedisForTest() {
	err := initConfigByViper()
	if err != nil {
		panic(err)
	}
	if err := initRedisDefault(); err != nil {
		panic(err)
	}
}

func initRedisDefault() error {
	return cache.InitRedisOpt()
}

func initConfigByViper() error {
	configFile := path.Join(pathCallUpperLevel(4), "conf", "config.yaml")
	fmt.Printf("=> configFile %v\n", configFile)
	// set timezone
	timezone.SetZoneUTC()
	// parse flag
	pflag.Parse()
	// init config
	if err := cfg.InitCfg(configFile, false); err != nil {
		fmt.Printf("Error, run service not use -c or config yaml error, more info: %v\n", err)
		return err
	}
	fmt.Printf("%s -> %v at time: %v\n", "run as name: ", viper.GetString("name"), time.Now().String())
	return nil
}

func pathCallUpperLevel(level int) string {
	_, fileName, _, _ := runtime.Caller(0)
	res := fileName
	for i := 0; i < level; i++ {
		dir := path.Dir(res)
		if dir == "" {
			return res
		}
		res = dir
	}
	return res
}
