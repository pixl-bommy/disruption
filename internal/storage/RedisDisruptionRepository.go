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
	// create disruption item
	disruptionItem, err := NewDisruptionEntry(name, description, userId)
	if err != nil {
		return "", err
	}
	disruptionItem.Status = DisruptionStatusActive

	// store disruption item in redis "event source"
	createdAt, err := r.storePartialSourceEntry(disruptionItem)
	if err != nil {
		return "", err
	}

	// update timestamps
	disruptionItem.CreatedAt = createdAt
	disruptionItem.ModifiedAt = createdAt

	// store disruption item in redis "derivate"
	err = r.storeDerivateEntry(disruptionItem)
	if err != nil {
		// TODO: If this happens, we should rebuild the derivations:disruptions
		// data structure from the "event source" stream.
		return "", err
	}

	// return the item id
	// NOTE: At this point we know UID is valid, as it was created by
	// NewDisruptionEntry. So we can safely parse it here.
	return disruptionItem.UID.String(), nil
}

func (r *RedisDisruptionRepository) Delete(uid, userId string) (bool, error) {
	// TODO: Maybe we should check, if there is a disruption with the given
	//       uid and userId. If not, we should return an error here.

	// create disruption item with only updated values
	disruptionItem, err := PartialDisruptionEntry(uid, userId)
	if err != nil {
		fmt.Println("Failed to create partial disruption entry:", err)
		return false, err
	}

	// set status to "deleted"
	disruptionItem.Status = DisruptionStatusDeleted

	// store disruption item in redis "event source"
	timestamp, err := r.storePartialSourceEntry(disruptionItem)
	if err != nil {
		fmt.Println("Failed to store partial disruption entry:", err)
		return false, err
	}

	// delete the item from the "derivate" store
	derivateKey := configs.KeyPrefix + ":" + disruptionsDerivateKey + ":" + uid
	err = r.client.Del(*r.context, derivateKey).Err()
	if err != nil {
		// TODO: If this happens, we should rebuild the derivations:disruptions
		fmt.Println("Failed to delete derivate entry:", err)
	}

	// update timestamp
	disruptionItem.ModifiedAt = timestamp

	return true, nil
}

func (r *RedisDisruptionRepository) GetAll() ([]DisruptionEntry, error) {
	return nil, errors.New("not implemented")
}

// Stores the a disruption item in the redis "event source" stream.
func (r *RedisDisruptionRepository) storePartialSourceEntry(disruptionItem *DisruptionEntry) (int64, error) {
	fmt.Println("storePartialSourceEntry:", disruptionItem)

	// create redis target key
	sourceKey := configs.KeyPrefix + ":" + disruptionsSourceKey

	// store disruption item as Stream in redis
	xAddArgs := &redis.XAddArgs{
		Stream: sourceKey, Values: disruptionItem.ToDisruptionEntryRaw(),
	}

	redisStreamId, err := r.client.XAdd(*r.context, xAddArgs).Result()
	if err != nil {
		return -1, err
	}

	// The stream id is defined as `{timestamp}-{sequence}` by default. As we use
	// the timestamp to determine creation and modification time, we can use the
	// the splitted stream id here.
	splitted := strings.Split(redisStreamId, "-")
	timestamp, err := strconv.ParseInt(splitted[0], 10, 64)
	if err != nil {
		// NOTE: If this happens, we have a problem with the redis stream id
		//       generation. This should never happen.
		fmt.Println("Failed to parse stream id:", err)
		return -1, err
	}

	return timestamp, nil
}

// Stores the a disruption item in the redis "derivate" stream.
func (r *RedisDisruptionRepository) storeDerivateEntry(disruptionItem *DisruptionEntry) error {
	// create redis target key
	derivateKey := configs.KeyPrefix + ":" + disruptionsDerivateKey + ":" + disruptionItem.UID.String()

	// store disruption item as Hash in redis
	err := r.client.HSet(*r.context, derivateKey, disruptionItem.ToDisruptionEntryRaw()).Err()
	if err != nil {
		return err
	}

	return nil
}
