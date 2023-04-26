package database

import (
	"context"
	"log"
	"sync"

	"github.com/NEXUS-CD/UPay/configs"
	"github.com/go-redis/redis/v8"
)

var (
	redisInstance *redis.Client
	redisOnce     sync.Once
)

func InitRedis() {
	redisOnce.Do(func() {
		conf := configs.GetConfig().Redis
		redisInstance = redis.NewClient(&redis.Options{
			Addr:     conf.Host + ":" + conf.Port,
			Password: conf.Password,
			DB:       conf.DB,
		})

		_, err := redisInstance.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}
	})
}

func GetRedisClient() *redis.Client {
	return redisInstance
}
