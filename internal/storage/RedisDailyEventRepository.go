package storage

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pixl-bommy/disruption/configs"
	"github.com/redis/go-redis/v9"
)

const (
	dailyEventSourceKey   = "sources:events:%s"
	dailyEventDerivateKey = "derivates:events:%s"
)

type RedisDailyEventRepository struct {
	client  *redis.Client
	context *context.Context
}

func NewRedisDailyEventRepository(client *redis.Client, context *context.Context) *RedisDailyEventRepository {
	return &RedisDailyEventRepository{client: client, context: context}
}

func (r *RedisDailyEventRepository) Create(disruption *DisruptionEntity, userId string) (string, error) {
	// 1. check user id is valid
	// 2. create daily event item
	// 3. store daily event item in redis "event source"
	// 4. update timestamps
	// 5. store daily event item in redis "derivate"
	// 6. return the item id

	err := uuid.Validate(userId)
	if err != nil {
		return "", err
	}

	dailyEvent, err := NewDailyEventEntity(disruption)
	if err != nil {
		return "", err
	}

	redisStreamId, err := r.storePartialSourceEntry(dailyEvent, userId)
	if err != nil {
		return "", err
	}

	dailyEvent.UID = redisStreamId

	err = r.storeDerivateEntry(dailyEvent, userId)
	if err != nil {
		return "", err
	}

	return dailyEvent.UID, nil
}

// Stores the a daily event item in the redis "event source" stream.
// Returns the created stream id, which is used as item id.
func (r *RedisDailyEventRepository) storePartialSourceEntry(dailyEvent *DailyEventEntity, userId string) (string, error) {
	fmt.Println("storePartialSourceEntry[DailyEventEntity]:", dailyEvent)

	// create redis target keys
	sourceKey := configs.KeyPrefix + ":" + fmt.Sprintf(dailyEventSourceKey, userId)

	// store disruption item as Stream in redis
	xAddArgs := &redis.XAddArgs{
		Stream: sourceKey, Values: dailyEvent.ToExportable(),
	}

	redisStreamId, err := r.client.XAdd(*r.context, xAddArgs).Result()
	if err != nil {
		return "", err
	}

	// return the created stream id
	return redisStreamId, nil
}

// Stores the a daily event item in the redis "event derivate" store.
func (r *RedisDailyEventRepository) storeDerivateEntry(dailyEvent *DailyEventEntity, userId string) error {
	fmt.Println("storeDerivateEntry[DailyEventEntity]:", dailyEvent)

	// create redis target keys
	derivateKey := configs.KeyPrefix + ":" + fmt.Sprintf(dailyEventDerivateKey, userId)

	// store disruption item as Hash in redis
	_, err := r.client.HSet(*r.context, derivateKey, dailyEvent.ToExportable()).Result()
	if err != nil {
		return err
	}

	// return the created stream id
	return nil
}
