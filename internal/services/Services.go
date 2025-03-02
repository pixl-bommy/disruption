package services

import "github.com/pixl-bommy/disruption/internal/storage"

var (
	Disruptions *DisruptionService
)

func init() {
	redisClient := storage.NewRedisClient()
	context := storage.NewContext()

	disruptionRepo := storage.NewRedisDisruptionRepository(redisClient, context)
	Disruptions = NewDisruptionService(disruptionRepo)
}
