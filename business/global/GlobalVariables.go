package global

import (
	"blogServe/business/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config      config.GlobalConfig
	DefaultDb   *gorm.DB
	RedisClient *redis.Client
)
