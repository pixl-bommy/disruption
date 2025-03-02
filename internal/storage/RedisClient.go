package storage

import (
	"context"
	"fmt"

	"github.com/pixl-bommy/disruption/configs"
	"github.com/redis/go-redis/v9"
)

// newRedisStorageManager creates a new redis storage manager.
func NewRedisClient() *redis.Client {
	redisAddress := configs.RedisHost + ":" + configs.RedisPort

	options := redis.Options{
		Addr:     redisAddress,
		DB:       configs.RedisDB,
		Username: configs.RedisUsername,
		Password: configs.RedisPassword,
	}
	redisClient := redis.NewClient(&options)

	fmt.Println("Created redis client.", options)

	return redisClient
}

func NewContext() *context.Context {
	context := context.Background()
	return &context
}
