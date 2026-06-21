package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	ctx         = context.Background()
)

func ConnectRedis() {

	RedisClient = redis.NewClient(
		&redis.Options{
			Addr: GetEnv("REDIS_HOST") + ":" + GetEnv("REDIS_PORT"),
		},
	)

	_, err := RedisClient.Ping(ctx).Result()

	if err != nil {
		log.Fatal("Failed connect Redis:", err)
	}

	log.Println("Redis Connected")
}
