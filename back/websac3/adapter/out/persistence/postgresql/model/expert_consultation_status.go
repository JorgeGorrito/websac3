package model

import "gorm.io/gorm"

type ExpertConsultationStatus struct {
	ID    uint                           `gorm:"primaryKey" json:"id"`
	Names []ExpertConsultationStatusName `json:"names"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (ExpertConsultationStatus) TableName() string {
	return "expert_consultation_statuses"
}
