package cache

import (
	"github.com/bridgewwater/golang-project-temple-docker-db/cfg"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/cache/optDefault"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/cache/optredis"
)

func InitRedisOpt() error {
	err := optredis.InitByConfigList(cfg.Global().RedisOptConfig)
	if err != nil {
		return err
	}

	return optDefault.Init()
}
