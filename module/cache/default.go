package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var defaultRedis *redis.Client
var ErrDefaultCacheEmpty = fmt.Errorf("default redis client is empty, must be init")
var ErrDefaultKeyEmpty = fmt.Errorf("default redis client key is empty plase check")
var ErrDefaultDataEmpty = fmt.Errorf("default redis client data is empty")

func Default() *redis.Client {
	if defaultRedis == nil {
		panic(fmt.Errorf("not use cache.InitRedisDefault() to init redis client"))
	}
	return defaultRedis
}

func InitRedisDefault() (*redis.Client, error) {
	redisByName, err := initRedisByName("default")
	if err != nil {
		return nil, err
	}
	defaultRedis = redisByName
	return defaultRedis, nil
}

func RedisDefaultExists(key string, prefix string) (bool, error) {
	if key == "" {
		return false, ErrDefaultKeyEmpty
	}
	return HasKey(defaultRedis, fmt.Sprintf("%v%v", prefix, key))
}

func RedisDefaultSet(key string, prefix string, data interface{}, expiration time.Duration) error {
	if key == "" {
		return ErrDefaultKeyEmpty
	}
	if data == nil {
		return fmt.Errorf("data is empty")
	}
	if defaultRedis == nil {
		return ErrDefaultCacheEmpty
	}
	marshal, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	return defaultRedis.
		Set(fmt.Sprintf(`%v%v`, prefix, key), string(marshal), expiration).
		Err()
}

func RedisDefaultGet(key string, prefix string, v interface{}) error {
	if key == "" {
		return ErrDefaultKeyEmpty
	}
	if defaultRedis == nil {
		return ErrDefaultCacheEmpty
	}
	result, err := defaultRedis.Get(fmt.Sprintf("%v%v", prefix, key)).Result()
	if err != nil {
		return err
	}
	if result == "" {
		return ErrDefaultDataEmpty
	}
	err = json.Unmarshal([]byte(result), &v)
	if err != nil {
		return err
	}
	return nil
}
