package optredis

import (
	"encoding/json"
	"fmt"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/cache/bloomfilter"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/cache/optredis/redisconfig"
	"github.com/go-redis/redis"
	"github.com/willf/bloom"
	"time"
)

type OptRedis struct {
	Name          string
	UseBoomFilter bool
	BloomK        uint
	BloomN        uint
	BloomM        uint

	RedisClient *redis.Client
	BloomFilter *bloom.BloomFilter

	errClientNotInit error
	errKeyEmpty      error
	errKeyFilter     error
	errDataEmpty     error
}

func (o OptRedis) Exists(key string, prefix string) (bool, error) {
	if key == "" {
		return false, o.errKeyEmpty
	}
	totalKey := fmt.Sprintf("%v%v", prefix, key)
	if o.UseBoomFilter {
		if o.BloomFilter.TestString(totalKey) {
			return ExistsKey(o.RedisClient, totalKey)
		} else {
			return false, nil
		}
	}
	return ExistsKey(o.RedisClient, totalKey)
}

func (o OptRedis) SetJson(key string, prefix string, data interface{}, expiration time.Duration) error {
	if key == "" {
		return o.errKeyEmpty
	}
	if data == nil {
		return o.errDataEmpty
	}
	if o.RedisClient == nil {
		return o.errClientNotInit
	}

	marshal, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	totalKey := fmt.Sprintf(`%v%v`, prefix, key)
	err = o.RedisClient.
		Set(totalKey, string(marshal), expiration).
		Err()
	if err != nil {
		return err
	}
	if o.UseBoomFilter {
		o.BloomFilter.AddString(totalKey)
	}
	return nil
}

func (o OptRedis) GetJson(key string, prefix string, v interface{}) error {
	if key == "" {
		return o.errKeyEmpty
	}
	if o.RedisClient == nil {
		return o.errClientNotInit
	}
	totalKey := fmt.Sprintf("%v%v", prefix, key)
	if o.UseBoomFilter {
		if !o.BloomFilter.TestString(totalKey) {
			return o.errKeyFilter
		}
	}
	result, err := o.RedisClient.Get(totalKey).Result()
	if err != nil {
		return err
	}
	if result == "" {
		return o.errDataEmpty
	}
	err = json.Unmarshal([]byte(result), &v)
	if err != nil {
		return err
	}
	return nil
}

func (o OptRedis) Client() *redis.Client {
	return o.RedisClient
}

func (o OptRedis) Ping() (OptRedis, error) {
	_, err := o.RedisClient.Ping().Result()
	if err != nil {
		return o, err
	}
	return o, nil
}

func (o OptRedis) InitByName() OptRedis {
	if redisConfigList == nil {
		panic(errRedisConfigListEmpty)
	}
	redisConf := redisconfig.ByName(*redisConfigList, o.Name)
	redisAddr := parseEnvStringOrDefault(parseEnvByName(o.Name, envKeyCacheRedisAddr), redisConf.Addr)
	redisPassword := parseEnvStringOrDefault(parseEnvByName(o.Name, envKeyCacheRedisPassword), redisConf.Password)
	redisDB := parseEnvIntOrDefault(parseEnvByName(o.Name, envKeyCacheRedisDB), redisConf.DB)
	redisMaxRetries := parseEnvIntOrDefault(parseEnvByName(o.Name, envKeyCacheRedisMaxRetries), redisConf.MaxRetries)
	redisDialTimeout := parseEnvIntOrDefault(parseEnvByName(o.Name, envKeyCacheRedisDialTimeout), redisConf.DialTimeout)
	redisReadTimeout := parseEnvIntOrDefault(parseEnvByName(o.Name, envKeyCacheRedisReadTimeout), redisConf.ReadTimeout)
	redisWriteTimeout := parseEnvIntOrDefault(parseEnvByName(o.Name, envKeyCacheRedisWriteTimeout), redisConf.WriteTimeout)

	o.RedisClient = redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		Password:     redisPassword, // no password set ""
		DB:           redisDB,       // use default DB 0
		MaxRetries:   redisMaxRetries,
		DialTimeout:  time.Duration(redisDialTimeout) * time.Second,
		ReadTimeout:  time.Duration(redisReadTimeout) * time.Second,
		WriteTimeout: time.Duration(redisWriteTimeout) * time.Second,
	})

	o.BloomFilter = bloomfilter.InitRedisFilter(5, 100, 20)

	return o
}

type RedisFunc interface {
	InitByName() OptRedis
	Ping() (OptRedis, error)
	Client() *redis.Client
	Exists(key string, prefix string) (bool, error)
	SetJson(key string, prefix string, data interface{}, expiration time.Duration) error
	GetJson(key string, prefix string, v interface{}) error
}

func NewOptRedis(cfg Config) RedisFunc {
	return &OptRedis{
		Name:          cfg.Name,
		UseBoomFilter: cfg.UseBoomFilter,

		errClientNotInit: fmt.Errorf("%v opt redis : client empty, must be init", cfg.Name),
		errKeyEmpty:      fmt.Errorf("%v opt redis : key is empty plase check", cfg.Name),
		errKeyFilter:     fmt.Errorf("%v opt redis : key filtered", cfg.Name),
		errDataEmpty:     fmt.Errorf("%v opt redis : data is empty", cfg.Name),
	}
}
