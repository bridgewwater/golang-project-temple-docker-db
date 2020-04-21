package cache

import (
	"encoding/json"
	"fmt"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/cache/bloomfilter"
	"github.com/go-redis/redis"
	"github.com/willf/bloom"
	"time"
)

var defaultRedis *redis.Client
var ErrDefaultCacheEmpty = fmt.Errorf("default redis client: is empty, must be init")
var ErrDefaultKeyEmpty = fmt.Errorf("default redis client: key is empty plase check")
var ErrDefaultKeyFilter = fmt.Errorf("default redis client: key filtered")
var ErrDefaultDataEmpty = fmt.Errorf("default redis client: data is empty")
var defaultRedisFilter *bloom.BloomFilter

func InitRedisDefault() (*redis.Client, error) {
	redisByName, err := initRedisByName("default")
	if err != nil {
		return nil, err
	}
	defaultRedis = redisByName
	defaultRedisFilter = bloomfilter.InitRedisFilter(5, 100, 20)
	return defaultRedis, nil
}

func RedisDefaultClient() *redis.Client {
	if defaultRedis == nil {
		panic(fmt.Errorf("not use cache.InitRedisDefault() to init redis client"))
	}
	return defaultRedis
}

type DefaultRedisBiz struct {
	useBoomFilter bool
	Client        *redis.Client
}

func RedisDefault(useBloom bool) *DefaultRedisBiz {
	d := &DefaultRedisBiz{
		useBoomFilter: useBloom,
		Client:        RedisDefaultClient(),
	}
	return d
}

func (c *DefaultRedisBiz) RedisDefaultExists(key string, prefix string) (bool, error) {
	if key == "" {
		return false, ErrDefaultKeyEmpty
	}
	totalKey := fmt.Sprintf("%v%v", prefix, key)
	if c.useBoomFilter {
		if defaultRedisFilter.TestString(totalKey) {
			return HasKey(c.Client, totalKey)
		} else {
			return false, nil
		}
	}
	return HasKey(c.Client, totalKey)
}

func (c *DefaultRedisBiz) RedisDefaultSet(key string, prefix string, data interface{}, expiration time.Duration) error {
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
	totalKey := fmt.Sprintf(`%v%v`, prefix, key)
	if c.useBoomFilter {
		defaultRedisFilter.AddString(totalKey)
	}
	return c.Client.
		Set(totalKey, string(marshal), expiration).
		Err()
}

func (c *DefaultRedisBiz) RedisDefaultGet(key string, prefix string, v interface{}) error {
	if key == "" {
		return ErrDefaultKeyEmpty
	}
	if defaultRedis == nil {
		return ErrDefaultCacheEmpty
	}
	totalKey := fmt.Sprintf("%v%v", prefix, key)
	if c.useBoomFilter {
		if !defaultRedisFilter.TestString(totalKey) {
			return ErrDefaultKeyFilter
		}
	}
	result, err := c.Client.Get(totalKey).Result()
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
