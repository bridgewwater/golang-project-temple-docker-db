package optredis

import (
	"fmt"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/cache/optredis/redisconfig"
)

var redisConfigList *[]redisconfig.RedisConfig

var errRedisConfigListEmpty = fmt.Errorf("optredis err: redis config list is empty")

func InitByConfigList(redisClientList []redisconfig.RedisConfig) error {
	if redisClientList == nil {
		return errRedisConfigListEmpty
	}
	if len(redisClientList) == 0 {
		return errRedisConfigListEmpty
	}
	redisConfigList = &redisClientList
	return nil
}
