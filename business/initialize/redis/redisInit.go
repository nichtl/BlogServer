package redis

import (
	"blogServe/business/global"
	"github.com/redis/go-redis/v9"
)

func InitRedis() (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.GetAddr(),
		Password: global.Config.Redis.Password, // no password set
		DB:       global.Config.Redis.DB,       // use default DB
	})
	return client
}
