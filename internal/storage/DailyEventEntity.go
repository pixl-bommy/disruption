package storage

import "github.com/google/uuid"

type DailyEventStatus string

const (
	DailyEventStatusDeleted  DailyEventStatus = "deleted"
	DailyEventStatusReplaced DailyEventStatus = "replaced"
)

type DailyEventEntity struct {
	UID            string           `json:"uid" redis:"uid"`
	DisruptionId   uuid.UUID        `json:"-" redis:"-"`
	DisruptionName string           `json:"-" redis:"-"`
	Status         DailyEventStatus `json:"-" redis:"-"`
	ReplacedBy     string           `json:"-" redis:"-"`
	Replaces       string           `json:"-" redis:"-"`

	IconName string `json:"-" redis:"-"`
	Color    string `json:"-" redis:"-"`
}

type DailyEventEntityExportable struct {
	UID            string `json:"uid"                      redis:"uid,omitempty"`
	DisruptionId   string `json:"disruptionId,omitempty"   redis:"disruption_id,omitempty"`
	DisruptionName string `json:"disruptionName,omitempty" redis:"disruption_name,omitempty"`
	Status         string `json:"status,omitempty"         redis:"status,omitempty"`
	ReplacedBy     string `json:"replacedBy,omitempty"     redis:"replaced_by,omitempty"`
	Replaces       string `json:"replaces,omitempty"       redis:"replaces,omitempty"`

	IconName string `json:"iconName,omitempty" redis:"icon_name,omitempty"`
	Color    string `json:"color,omitempty"    redis:"color,omitempty"`
}

func NewDailyEventEntity(disruption *DisruptionEntity) (*DailyEventEntity, error) {
	return &DailyEventEntity{
		DisruptionId:   disruption.UID,
		DisruptionName: disruption.Name,
	}, nil
}

func (ee *DailyEventEntityExportable) ToEntity() DailyEventEntity {
	return DailyEventEntity{
		UID:            ee.UID,
		DisruptionId:   uuid.MustParse(ee.DisruptionId),
		DisruptionName: ee.DisruptionName,
		Status:         DailyEventStatus(ee.Status),
		ReplacedBy:     ee.ReplacedBy,
		Replaces:       ee.Replaces,

		IconName: ee.IconName,
		Color:    ee.Color,
	}
}

func (ee *DailyEventEntity) ToExportable() DailyEventEntityExportable {
	return DailyEventEntityExportable{
		UID:            ee.UID,
		DisruptionId:   ee.DisruptionId.String(),
		DisruptionName: ee.DisruptionName,
		Status:         string(ee.Status),
		ReplacedBy:     ee.ReplacedBy,
		Replaces:       ee.Replaces,

		IconName: ee.IconName,
		Color:    ee.Color,
	}
}
