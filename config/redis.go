package config

import (
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

var Redis = redis.Client{}

func RedisInit() *redis.Client {
	REDIS_HOST := "localhost"
	REDIS_PORT := "6379"
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

	RedisClient := redis.NewClient(&redis.Options{
		Addr: REDIS_HOST + ":" + REDIS_PORT,
		DB:   REDIS_DB,
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		panic("redis ping error")
	}
	Redis = *RedisClient
	return RedisClient
}
