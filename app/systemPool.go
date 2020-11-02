package funnel

import "github.com/go-redis/redis"

var Redis = RedisInit()

func RedisInit() *redis.Client {
	RedisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}
	return RedisClient
}
