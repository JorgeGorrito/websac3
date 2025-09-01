package model

import "gorm.io/gorm"

type Topic struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	Names []TopicName

	KnowledgeAreaID uint          `gorm:"not null" json:"ka_id"`
	KnowledgeArea   KnowledgeArea `gorm:"foreignKey:KnowledgeAreaID"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
