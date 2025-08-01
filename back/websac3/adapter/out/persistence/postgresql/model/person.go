package model

import "time"

type Person struct {
	ID uint `gorm:"primaryKey" `

	Name                 string `gorm:"type:varchar(36); not null" `
	Lastname             string `gorm:"type:varchar(36); not null" `
	IdentificationNumber string `gorm:"type:varchar(16); not null; unique" json:"identification_number"`

	IdentificationTypeID uint `gorm:"not null" json:"identification_type_id"`
	IdentificationType   IdentificationType

	HigherEducationInstitutionSnies uint                       `gorm:"not null" json:"higher_education_institution_snies"`
	HigherEducationInstitution      HigherEducationInstitution `gorm:"foreignKey:HigherEducationInstitutionSnies"`

	JobPosition string `gorm:"type:varchar(128); not null" `

	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeleteAt  *time.Time `gorm:"null; default:null" `
}
