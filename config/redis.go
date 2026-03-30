package config

import (
	"exchangeapp/global"
	"log"

	"github.com/go-redis/redis"
)

// 初始化redis
func initRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis,got error:%v", err)
	}
	global.RedisDB = RedisClient

}
