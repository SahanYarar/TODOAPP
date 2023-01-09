package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	rdb *redis.Client
}

type RedisRepositoryInterface interface {
	Set(ctx context.Context, userId string, token string, exp time.Duration) error
	Get(ctx context.Context, userId string) string
	Delete(ctx context.Context, userId string) error
	Exists(ctx context.Context, userId string) int64
}

func CreateRepositoryRedis(redisclient *redis.Client) *RedisClient {
	return &RedisClient{rdb: redisclient}

}

func (redisClient *RedisClient) Set(ctx context.Context, userId string, token string, exp time.Duration) error {

	return redisClient.rdb.Set(ctx, userId, token, exp).Err()

}

func (redisClient *RedisClient) Get(ctx context.Context, userId string) string {

	result, err := redisClient.rdb.Get(ctx, userId).Result()
	if err != nil {
		panic(err)
	}

	return result

}

func (redisClient *RedisClient) Delete(ctx context.Context, userId string) error {

	err := redisClient.rdb.Del(ctx, userId).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func (redisclient *RedisClient) Exists(ctx context.Context, userId string) int64 {
	result, err := redisclient.rdb.Exists(ctx, userId).Result()
	if err != nil {
		panic(err)
	}
	return result
}
