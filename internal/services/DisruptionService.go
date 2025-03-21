package services

import "github.com/pixl-bommy/disruption/internal/storage"

type DisruptionService struct {
	repo storage.DisruptionRepository
}

func NewDisruptionService(repo storage.DisruptionRepository) *DisruptionService {
	return &DisruptionService{repo: repo}
}

func (s *DisruptionService) Create(name, description, userId string) (string, error) {
	return s.repo.Create(name, description, userId)
}

func (s *DisruptionService) Delete(uid, userId string) (bool, error) {
	return s.repo.Delete(uid, userId)
}

func (s *DisruptionService) Get(uid string) (*storage.DisruptionEntity, error) {
	return s.repo.Get(uid)
}

func (s *DisruptionService) GetAll() ([]*storage.DisruptionEntityExportable, error) {
	return s.repo.GetAll()
}
