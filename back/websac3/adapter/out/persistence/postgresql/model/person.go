package model

import "time"

type Person struct {
	ID uint `gorm:"primaryKey" `

	Name                 string `gorm:"type:varchar(36); not null" `
	Lastname             string `gorm:"type:varchar(36); not null" `
	IdentificationNumber string `gorm:"type:varchar(16); not null; unique" `

	IdentificationTypeID uint `gorm:"not null" `
	IdentificationType   IdentificationType

	HigherEducationInstitutionSnies uint                       `gorm:"not null" `
	HigherEducationInstitution      HigherEducationInstitution `gorm:"foreignKey:HigherEducationInstitutionSnies"`

	JobPosition string `gorm:"type:varchar(128); not null" `

	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time `gorm:"null; default:null" `
}
