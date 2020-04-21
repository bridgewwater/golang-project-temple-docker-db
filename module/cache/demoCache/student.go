package demoCache

import (
	"github.com/bridgewwater/golang-project-temple-docker-db/module/cache"
	"github.com/bridgewwater/golang-project-temple-docker-db/module/demo"
	"time"
)

const (
	cpStudentPrefix string = "cache-student-"
	// one week
	cpStudentExpiration = time.Duration(24) * time.Hour
)

func ExistsStudent(key string) (bool, error) {
	return cache.RedisDefault(false).
		RedisDefaultExists(key, cpStudentPrefix)
}

func Set(key string, data *demo.Student) error {
	return cache.RedisDefault(false).
		RedisDefaultSet(key, cpStudentPrefix, data, cpStudentExpiration)
}

func Get(key string, data *demo.Student) error {
	return cache.RedisDefault(false).
		RedisDefaultGet(key, cpStudentPrefix, data)
}
