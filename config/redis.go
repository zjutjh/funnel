package config

import (
	redisConfig "funnel/config/redis"
	"github.com/go-redis/redis"
)

var Redis *redis.Client
var RedisInfo redisConfig.RedisConfig

func init() {
	info := redisConfig.GetConfig()

	Redis = redis.NewClient(&redis.Options{
		Addr:     info.Host + ":" + info.Port,
		Password: info.Password,
		DB:       info.DB,
	})
	RedisInfo = info
}
