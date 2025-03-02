package storage

import "github.com/google/uuid"

type DisruptionStatus string

const (
	DisruptionStatusActive  DisruptionStatus = "active"
	DisruptionStatusDeleted DisruptionStatus = "deleted"
)

type DisruptionEntry struct {
	UID         uuid.UUID        `json:"-" redis:"-"`
	Name        string           `json:"-" redis:"-"`
	Description string           `json:"-" redis:"-"`
	Status      DisruptionStatus `json:"-" redis:"-"`
	ModifiedBy  uuid.UUID        `json:"-" redis:"-"`
	CreatedAt   int64            `json:"-" redis:"-"`
	ModifiedAt  int64            `json:"-" redis:"-"`
}

type DisruptionEntryRaw struct {
	UID         string `json:"uid"                   redis:"uid"`
	Name        string `json:"name,omitempty"        redis:"name,omitempty"`
	Description string `json:"description,omitempty" redis:"description,omitempty"`
	Status      string `json:"status,omitempty"      redis:"status,omitempty"`
	ModifiedBy  string `json:"modifiedBy,omitempty"  redis:"modified_by,omitempty"`
	CreatedAt   int64  `json:"createdAt,omitempty"   redis:"-"`
	ModifiedAt  int64  `json:"modifiedAt,omitempty"  redis:"-"`
}

func NewDisruptionEntry(name, description string, modifiedBy string) (*DisruptionEntry, error) {
	modifierUserId, err := uuid.Parse(modifiedBy)
	if err != nil {
		return nil, err
	}

	return &DisruptionEntry{
		UID:         uuid.New(),
		Name:        name,
		Description: description,
		Status:      DisruptionStatusActive,
		ModifiedBy:  modifierUserId,
	}, nil
}

func (d *DisruptionEntryRaw) ToDisruptionEntry() DisruptionEntry {
	return DisruptionEntry{
		UID:         uuid.MustParse(d.UID),
		Name:        d.Name,
		Description: d.Description,
		Status:      DisruptionStatus(d.Status),
		ModifiedBy:  uuid.MustParse(d.ModifiedBy),
		CreatedAt:   d.CreatedAt,
		ModifiedAt:  d.ModifiedAt,
	}
}

func (d *DisruptionEntry) ToDisruptionEntryRaw() DisruptionEntryRaw {
	return DisruptionEntryRaw{
		UID:         d.UID.String(),
		Name:        d.Name,
		Description: d.Description,
		Status:      string(d.Status),
		ModifiedBy:  d.ModifiedBy.String(),
		CreatedAt:   d.CreatedAt,
		ModifiedAt:  d.ModifiedAt,
	}
}
