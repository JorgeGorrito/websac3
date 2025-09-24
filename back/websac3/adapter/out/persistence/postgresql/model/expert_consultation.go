package model

import (
	"time"

	"gorm.io/gorm"
)

type ExpertConsultation struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Usuario que solicita la asesoría
	RequesterID uint `gorm:"not null;index" json:"requester_id"`
	Requester   User `gorm:"foreignKey:RequesterID"`

	// Programa de grado asociado
	DegreeProgramID uint          `gorm:"not null;index" json:"degree_program_id"`
	DegreeProgram   DegreeProgram `gorm:"foreignKey:DegreeProgramID"`

	// Reporte asociado
	ReportID uint   `gorm:"not null;index" json:"report_id"`
	Report   Report `gorm:"foreignKey:ReportID"`

	// Mensaje de solicitud (opcional)
	RequestMessage *string `gorm:"type:text" json:"request_message"`

	// Respuesta del experto
	ExpertResponse *string `gorm:"type:text" json:"expert_response"`

	// Usuario experto que proporciona la respuesta
	ExpertID *uint `gorm:"index" json:"expert_id"`
	Expert   *User `gorm:"foreignKey:ExpertID"`

	// Estado de la consulta
	StatusID uint                     `gorm:"not null;index;default:1" json:"status_id"`
	Status   ExpertConsultationStatus `gorm:"foreignKey:StatusID" json:"status"`

	// Fechas de auditoría
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	AnsweredAt *time.Time     `gorm:"default:null" json:"answered_at"`
	ClosedAt   *time.Time     `gorm:"default:null" json:"closed_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (ExpertConsultation) TableName() string {
	return "expert_consultations"
}
