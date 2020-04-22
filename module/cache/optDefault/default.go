package optDefault

import "github.com/bridgewwater/golang-project-temple-docker-db/module/cache/optredis"

const (
	name string = "default"
)

var defaultOptRedis *optredis.OptRedis

func Init() error {
	if defaultOptRedis == nil {
		config := optredis.NewConfig(
			optredis.WithName("default"),
			optredis.WithUseBoomFilter(true),
			optredis.WithUseBloomK(20),
			optredis.WithUseBloomN(1000),
			optredis.WithUseBloomM(5),
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
