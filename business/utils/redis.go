package utils

import (
	"blogServe/business/global"
	"context"
	"time"
)

type redis struct{}

var (
	localCache = NewLocalCache()
)

func Set(key string, value string, timeVal int, duration time.Duration) error {
	redisClient := global.RedisClient
	expiration := time.Duration(timeVal) * duration
	if redisClient != nil {
		return redisClient.Set(context.Background(), key, value, expiration).Err()
	}
	localCache.Set(key, value, expiration)
	return nil
}

func Get(key string) (string, error) {
	redisClient := global.RedisClient
	if redisClient != nil {
		return redisClient.Get(context.Background(), key).Result()
	}

	val, succ := localCache.Get(key)
	if succ {
		return val, nil
	}
	return "", nil
}

func Del(key string) error {
	redisClient := global.RedisClient
	if redisClient != nil {
		return redisClient.Del(context.Background(), key).Err()
	}
	localCache.Delete(key)
	return nil
}
