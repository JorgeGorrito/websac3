package model

import "gorm.io/gorm"

type KnowledgeArea struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	Names []KnowledgeAreaName

	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
