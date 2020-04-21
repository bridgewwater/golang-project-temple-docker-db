package bloomfilter

import (
	"fmt"
	"github.com/willf/bloom"
)

var redisFilter *bloom.BloomFilter

func AddRedis(data []byte) {
	Redis().Add(data)
}

func RedisTest(data []byte) bool {
	return Redis().Test(data)
}

func RedisTestAndAdd(data []byte) bool {
	return Redis().TestAndAdd(data)
}

func Redis() *bloom.BloomFilter {
	if redisFilter == nil {
		panic(fmt.Errorf("not use bloomfilter.InitRedisFilter() to init"))
	}
	return redisFilter
}

// init redis filter, like of 20, 5 keys (k, m) in 1000 (n)
//	k -> the number of hashing functions on elements of the set, The actual hashing functions are important, too, but this is not a parameter for this implementation
//	n -> data size
//	m -> maximum size, typically a reasonably large multiple of the cardinality of the set to represent
func InitRedisFilter(k, n, m uint) *bloom.BloomFilter {
	if redisFilter == nil {
		redisFilter = bloom.New(k*n, m)
	}
	return redisFilter
}
