package common

import (
	"github.com/go-redis/redis/v8"
)

func CreateRedisClient() *redis.Client {
	env := GetEnvironment()
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.RedisAddr,
		Password: env.RedisPassword,
		DB:       0,
	})
	return rdb
}
