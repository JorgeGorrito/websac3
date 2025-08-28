package model

import "time"

type KnowledgeArea struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	Names []KnowledgeAreaName

	DeletedAt *time.Time `gorm:"index; default: null;"`
}
