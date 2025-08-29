package model

import "time"

type DurationUnit struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	Names []DurationUnitName

	DeletedAt *time.Time `gorm:"index; default: null;"`
}
