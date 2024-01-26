package config

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var Redis = redis.Client{}

func RedisInit() *redis.Client {
	REDIS_HOST := "localhost"
	REDIS_PORT := "6379"
	REDIS_PASSWORD := ""
	REDIS_DB := 0

	if os.Getenv("REDIS_HOST") != "" {
		REDIS_HOST = os.Getenv("REDIS_HOST")
	}

	if os.Getenv("REDIS_PORT") != "" {
		REDIS_PORT = os.Getenv("REDIS_PORT")
	}

	if os.Getenv("REDIS_DB") != "" {
		REDIS_DB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	}

	if os.Getenv("REDIS_PASSWORD") != "" {
		REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	}

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST + ":" + REDIS_PORT,
		Password: REDIS_PASSWORD,
		DB:       REDIS_DB,
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		panic("redis ping error")
	}
	Redis = *RedisClient
	return RedisClient
}
