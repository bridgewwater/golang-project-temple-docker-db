package main

import (
	"fmt"
	"github.com/bridgewwater/golang-project-temple-docker-db/cfg"
	"github.com/sinlovgo/timezone"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
)

var (
	pFlagCLI = pflag.StringP("config", "c", "", "api server config file path.")
)

func main() {
	// set timezone
	timezone.SetZoneUTC()
	// parse flag
	pflag.Parse()
	// init config
	if err := cfg.Init(*pFlagCLI); err != nil {
		fmt.Printf("Error, run service not use -c or config yaml error, more info: %v\n", err)
		panic(err)
	}
	fmt.Printf("%s -> %v at time: %v\n", "run as name: ", viper.GetString("name"), time.Now().String())
}
