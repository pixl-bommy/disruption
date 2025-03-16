package storage

type DisruptionRepository interface {
	Create(name, description, userId string) (string, error)
	Delete(uid, userId string) (bool, error)
	GetAll() ([]*DisruptionEntityExportable, error)
}
