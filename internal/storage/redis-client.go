package storage

import (
	"context"
	"fmt"

	"github.com/pixl-bommy/disruption/configs"
	"github.com/redis/go-redis/v9"
)

// redisManager is the manager for all redis operations.
type redisManager struct {
	*redis.Client
}

var (
	singleInstance *redisManager

	// redisContext is the context for all redis operations.
	redisContext *context.Context
)

func Redis() *redisManager {
	if singleInstance == nil {
		fmt.Println("Creating single instance now.")
		singleInstance = newRedisStorageManager()
	}

	return singleInstance
}

func GetContext() *context.Context {
	if redisContext == nil {
		fmt.Println("Creating single context now.")

		newContext := context.Background()
		redisContext = &newContext
	}

	return redisContext
}

// newRedisStorageManager creates a new redis storage manager.
func newRedisStorageManager() *redisManager {
	redisAddress := configs.RedisHost + ":" + configs.RedisPort

	options := redis.Options{
		Addr:     redisAddress,
		DB:       configs.RedisDB,
		Username: configs.RedisUsername,
		Password: configs.RedisPassword,
	}

	redisClient := redis.NewClient(&options)

	return &redisManager{
		redisClient,
	}
}
