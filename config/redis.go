package config

import (
	redisConfig "funnel/config/redis"
	"time"

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

	Redis.Set("CONNECTION_TEST", "CONNECTION_TEST", time.Second)
	redisVal, err := Redis.Get("CONNECTION_TEST").Result()
	if err != nil || redisVal != "CONNECTION_TEST" {
		panic("Redis connection failed")
	}
}
