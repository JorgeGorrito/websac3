package model

import "time"

type Topic struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	Names []TopicName

	KnowledgeAreaID uint          `gorm:"not null" json:"ka_id"`
	KnowledgeArea   KnowledgeArea `gorm:"foreignKey:KnowledgeAreaID"`

	DeletedAt *time.Time `gorm:"index; default: null;"`
}
