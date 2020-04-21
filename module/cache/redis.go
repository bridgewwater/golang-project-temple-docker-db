package cache

import (
	"fmt"
	"github.com/bridgewwater/golang-project-temple-docker-db/cfg"
	"github.com/bridgewwater/golang-project-temple-docker-db/pkg/zlog"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"strings"
	"time"
)

const (
	envKeyCacheRedisPrefix       = "CACHE_REDIS"
	envKeyCacheRedisAddr         = "ADDR"
	envKeyCacheRedisPassword     = "PASSWORD"
	envKeyCacheRedisDB           = "DB"
	envKeyCacheRedisMaxRetries   = "MAX_RETRIES"
	envKeyCacheRedisDialTimeout  = "DIAL_TIMEOUT"
	envKeyCacheRedisReadTimeout  = "READ_TIMEOUT"
	envKeyCacheRedisWriteTimeout = "WRITE_TIMEOUT"
)

var ErrCacheStatus = fmt.Errorf("cache status error")

func parseEnvStringOrDefault(envKey, defaultStr string) string {
	var result string
	if viper.GetString(envKey) == "" {
		result = defaultStr
	} else {
		result = viper.GetString(envKey)
	}
	return result
}

func parseEnvIntOrDefault(envKey string, defaultInt int) int {
	var result int
	if viper.GetInt(envKey) == 0 {
		result = defaultInt
	} else {
		result = viper.GetInt(envKey)
	}
	return result
}

// in viper yaml
//	redis_clients:
//  - name: default
//    addr: localhost:35021
//    password:
//    db: 0
//    max_retries: 0 # Default is to not retry failed commands
//    dial_timeout: 5 # Default is 5 seconds.
//    read_timeout: 3 # Default is 3 seconds.
//    write_timeout: 3 # Default is ReadTimeout
// and will cover by ENV name like DEFAULT
//	ENV_WEB_CACHE_REDIS_DEFAULT_ADDR
//	ENV_WEB_CACHE_REDIS_DEFAULT_PASSWORD
//	ENV_WEB_CACHE_REDIS_DEFAULT_DB
//	ENV_WEB_CACHE_REDIS_DEFAULT_MAX_RETRIES
//	ENV_WEB_CACHE_REDIS_DEFAULT_DIAL_TIMEOUT
//	ENV_WEB_CACHE_REDIS_DEFAULT_READ_TIMEOUT
//	ENV_WEB_CACHE_REDIS_DEFAULT_WRITE_TIMEOUT
func initRedisByName(name string) (*redis.Client, error) {
	redisConf := cfg.RedisClientByName(name)
	redisAddr := parseEnvStringOrDefault(parseEnvByName(name, envKeyCacheRedisAddr), redisConf.Addr)
	redisPassword := parseEnvStringOrDefault(parseEnvByName(name, envKeyCacheRedisPassword), redisConf.Password)
	redisDB := parseEnvIntOrDefault(parseEnvByName(name, envKeyCacheRedisDB), redisConf.DB)
	redisMaxRetries := parseEnvIntOrDefault(parseEnvByName(name, envKeyCacheRedisMaxRetries), redisConf.MaxRetries)
	redisDialTimeout := parseEnvIntOrDefault(parseEnvByName(name, envKeyCacheRedisDialTimeout), redisConf.DialTimeout)
	redisReadTimeout := parseEnvIntOrDefault(parseEnvByName(name, envKeyCacheRedisReadTimeout), redisConf.ReadTimeout)
	redisWriteTimeout := parseEnvIntOrDefault(parseEnvByName(name, envKeyCacheRedisWriteTimeout), redisConf.WriteTimeout)

	client := redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		Password:     redisPassword, // no password set
		DB:           redisDB,       // use default DB
		MaxRetries:   redisMaxRetries,
		DialTimeout:  time.Duration(redisDialTimeout) * time.Second,
		ReadTimeout:  time.Duration(redisReadTimeout) * time.Second,
		WriteTimeout: time.Duration(redisWriteTimeout) * time.Second,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		zlog.S().Errorf("redis conn pong: %v , err: %v", pong, err)
		return nil, err
	} else {
		zlog.S().Infof("redis conn success pong: %v", pong)
	}
	return client, nil
}

func parseEnvByName(name string, last string) string {
	return fmt.Sprintf("%s_%s_%s", envKeyCacheRedisPrefix, strings.ToUpper(name), last)
}

func HasKey(redisCli *redis.Client, key string) (bool, error) {
	if redisCli == nil {
		return false, ErrCacheStatus
	}
	count, err := redisCli.Exists(key).Result()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// scan keys instead redisCli.Keys()
//	redisCli *redis.Client
//	match string
//	maxCount int64
// return
//	error scan error
//	[]string removed repeated key
func RedisScanKeysMatch(redisCli *redis.Client, match string, maxCount int64) (error, []string) {
	var cursor uint64
	var scanFull []string
	for {
		keys, cursor, err := redisCli.Scan(cursor, match, maxCount).Result()
		if err != nil {
			return err, nil
		}
		if len(keys) > 0 {
			for _, v := range keys {
				scanFull = append(scanFull, v)
			}
		}
		if cursor == 0 {
			break
		}
	}
	scanRes := removeRepeatedElementString(scanFull)
	return nil, scanRes
}

func removeRepeatedElementString(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
