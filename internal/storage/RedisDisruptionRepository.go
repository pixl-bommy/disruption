package storage

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/pixl-bommy/disruption/configs"
	"github.com/redis/go-redis/v9"
)

const (
	disruptionsSourceKey   = "sources:disruptions"
	disruptionsDerivateKey = "derivates:disruptions"
)

type RedisDisruptionRepository struct {
	client  *redis.Client
	context *context.Context
}

func NewRedisDisruptionRepository(client *redis.Client, context *context.Context) *RedisDisruptionRepository {
	return &RedisDisruptionRepository{client: client, context: context}
}

func (r *RedisDisruptionRepository) Create(name, description, userId string) (string, error) {

	// create redis target key
	sourceKey := configs.KeyPrefix + ":" + disruptionsSourceKey

	// create disruption item
	disruptionItem, err := NewDisruptionEntry(name, description, userId)
	if err != nil {
		return "", err
	}

	disruptionItem.Status = DisruptionStatusActive

	// store disruption item in redis
	xAddArgs := &redis.XAddArgs{
		Stream: sourceKey, Values: disruptionItem.ToDisruptionEntryRaw(),
	}
	redisStreamId, err := r.client.XAdd(*r.context, xAddArgs).Result()
	if err != nil {
		return "", err
	}

	// TODO: This is not needed at this point, as we return the item id only.
	// The stream id is defined as `{timestamp}-{sequence}` by default. As we use
	// the timestamp to determine creation and modification time, we can use the
	// the splitted stream id here.
	splitted := strings.Split(redisStreamId, "-")
	createdAt, err := strconv.ParseInt(splitted[0], 10, 64)
	if err != nil {
		// NOTE: If this happens, we have a problem with the redis stream id
		//       generation. This should never happen.
		fmt.Println("Failed to parse stream id:", err)
		return "", err
	}

	// update the created at and modified at fields
	disruptionItem.CreatedAt = createdAt
	disruptionItem.ModifiedAt = createdAt

	fmt.Println("Create: ")
	fmt.Println("   ", sourceKey)
	fmt.Println("   ", redisStreamId)
	fmt.Println("   ", disruptionItem)

	// return the item id
	// NOTE: At this point we know UID is valid, as it was created by
	// NewDisruptionEntry. So we can safely parse it here.
	return disruptionItem.UID.String(), nil
}

func (r *RedisDisruptionRepository) Delete(uid, userId string) error {
	return errors.New("not implemented")
}

func (r *RedisDisruptionRepository) GetAll() ([]DisruptionEntry, error) {
	return nil, errors.New("not implemented")
}
