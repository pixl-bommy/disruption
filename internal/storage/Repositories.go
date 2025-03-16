package storage

type DisruptionRepository interface {
	Create(name, description, userId string) (string, error)
	Delete(uid, userId string) (bool, error)
	Get(uid string) (*DisruptionEntity, error)
	GetAll() ([]*DisruptionEntityExportable, error)
}

type DailyEventRepository interface {
	Create(disruption *DisruptionEntity, userId string) (string, error)
	Get(from, to int64) ([]*DailyEventEntityExportable, error)
}
