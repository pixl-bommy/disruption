package services

import "github.com/pixl-bommy/disruption/internal/storage"

type DailyEventService struct {
	repo storage.DailyEventRepository
}

func NewDailyEventService(repo storage.DailyEventRepository) *DailyEventService {
	return &DailyEventService{repo: repo}
}

func (s *DailyEventService) Create(disruptionId string, userId string) (string, error) {
	disruption, err := Disruptions.Get(disruptionId)
	if err != nil {
		return "", err
	}

	return s.repo.Create(disruption, userId)
}
