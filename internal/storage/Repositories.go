package storage

type DisruptionRepository interface {
	Create(name, description, userId string) (string, error)
	Delete(uid, userId string) error
	GetAll() ([]DisruptionEntry, error)
}
