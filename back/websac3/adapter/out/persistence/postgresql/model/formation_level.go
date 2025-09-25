package model

import "gorm.io/gorm"

type FormationLevel struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	Names []FormationLevelName

	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
