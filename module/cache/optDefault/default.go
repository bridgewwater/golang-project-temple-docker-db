package optDefault

import "github.com/sinlovgo/optredis"

const (
	name string = "default"
)

var defaultOptRedis *optredis.OptRedis

func Init() error {
	if defaultOptRedis == nil {
		config := optredis.NewConfig(
			optredis.WithName("default"),
		)
		optRedis, err := optredis.NewOptRedis(*config).InitByName().Ping()
		if err != nil {
			return err
		}
		defaultOptRedis = &optRedis
	}
	return nil
}

func Opt() *optredis.OptRedis {
	return defaultOptRedis
}
