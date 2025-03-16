package storage

import "github.com/google/uuid"

type DisruptionStatus string

const (
	DisruptionStatusActive  DisruptionStatus = "active"
	DisruptionStatusDeleted DisruptionStatus = "deleted"
)

type DisruptionEntity struct {
	UID         uuid.UUID        `json:"-" redis:"-"`
	Name        string           `json:"-" redis:"-"`
	Description string           `json:"-" redis:"-"`
	Status      DisruptionStatus `json:"-" redis:"-"`
	ModifiedBy  uuid.UUID        `json:"-" redis:"-"`
	CreatedAt   int64            `json:"-" redis:"-"`
	ModifiedAt  int64            `json:"-" redis:"-"`
}

type DisruptionEntityExportable struct {
	UID         string `json:"uid"                   redis:"uid"`
	Name        string `json:"name,omitempty"        redis:"name,omitempty"`
	Description string `json:"description,omitempty" redis:"description,omitempty"`
	Status      string `json:"status,omitempty"      redis:"status,omitempty"`
	ModifiedBy  string `json:"modifiedBy,omitempty"  redis:"modified_by,omitempty"`
	CreatedAt   int64  `json:"createdAt,omitempty"   redis:"created_at,omitempty"`
	ModifiedAt  int64  `json:"modifiedAt,omitempty"  redis:"modified_at,omitempty"`
}

func NewDisruptionEntity(name, description string, modifiedBy string) (*DisruptionEntity, error) {
	modifierUserId, err := uuid.Parse(modifiedBy)
	if err != nil {
		return nil, err
	}

	return &DisruptionEntity{
		UID:         uuid.New(),
		Name:        name,
		Description: description,
		Status:      DisruptionStatusActive,
		ModifiedBy:  modifierUserId,
	}, nil
}

func PartialDisruptionEntity(uid, modifiedBy string) (*DisruptionEntity, error) {
	uidUUID, err := uuid.Parse(uid)
	if err != nil {
		return nil, err
	}

	modifierUserId, err := uuid.Parse(modifiedBy)
	if err != nil {
		return nil, err
	}

	return &DisruptionEntity{
		UID:        uidUUID,
		ModifiedBy: modifierUserId,
	}, nil
}

func (d *DisruptionEntityExportable) ToEntity() DisruptionEntity {
	return DisruptionEntity{
		UID:         uuid.MustParse(d.UID),
		Name:        d.Name,
		Description: d.Description,
		Status:      DisruptionStatus(d.Status),
		ModifiedBy:  uuid.MustParse(d.ModifiedBy),
		CreatedAt:   d.CreatedAt,
		ModifiedAt:  d.ModifiedAt,
	}
}

func (d *DisruptionEntity) ToExportable() DisruptionEntityExportable {
	return DisruptionEntityExportable{
		UID:         d.UID.String(),
		Name:        d.Name,
		Description: d.Description,
		Status:      string(d.Status),
		ModifiedBy:  d.ModifiedBy.String(),
		CreatedAt:   d.CreatedAt,
		ModifiedAt:  d.ModifiedAt,
	}
}
