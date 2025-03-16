package storage

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pixl-bommy/disruption/configs"
	"github.com/redis/go-redis/v9"
)

const (
	dailyEventSourceKey = "sources:events:%s"
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

	dailyEventId, err := r.storePartialSourceEntry(dailyEvent, userId)
	if err != nil {
		return "", err
	}

	return dailyEventId, nil
}

func (r *RedisDailyEventRepository) Get(from, to int64) ([]*DailyEventEntityExportable, error) {
	// 1. check from and to are valid
	// 2. get daily event items from redis "event derivate"
	// 3. return the items

	if from < 0 || to < 0 || from > to {
		return nil, fmt.Errorf("invalid from/to timestamps")
	}

	// get all daily event items from redis
	derivateKey := configs.KeyPrefix + ":" + fmt.Sprintf(dailyEventSourceKey, configs.OneAndOnlyUserUid)
	derivateItems, err := r.client.XRange(*r.context, derivateKey, fmt.Sprintf("%d", from), fmt.Sprintf("%d", to)).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println("Get[DailyEventEntity]:", derivateItems)

	// convert the items to daily event entities
	dailyEvents := make([]*DailyEventEntityExportable, 0)
	for _, item := range derivateItems {
		dailyEvent := DailyEventEntityExportable{UID: item.ID}

		if disruptionId, ok := item.Values["disruption_id"].(string); ok {
			dailyEvent.DisruptionId = disruptionId
		}

		if disruptionName, ok := item.Values["disruption_name"].(string); ok {
			dailyEvent.DisruptionName = disruptionName
		}

		if status, ok := item.Values["status"].(string); ok {
			dailyEvent.Status = status
		}

		if iconName, ok := item.Values["icon_name"].(string); ok {
			dailyEvent.IconName = iconName
		}

		if color, ok := item.Values["color"].(string); ok {
			dailyEvent.Color = color
		}

		dailyEvents = append(dailyEvents, &dailyEvent)
	}

	// return the daily event entities
	return dailyEvents, nil
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
