package models

import "time"

type Person struct {
	ID uint `gorm:"primaryKey" mapper:"personID"`

	Name                 string `gorm:"type:varchar(36); not null" mapper:"personName"`
	Lastname             string `gorm:"type:varchar(36); not null" mapper:"personLastname"`
	IdentificationNumber string `gorm:"type:varchar(16); not null; unique" mapper:"personIdentificationNumber"`

	IdentificationTypeID uint `gorm:"not null" mapper:"personIdentificationTypeID"`
	IdentificationType   IdentificationType

	HigherEducationInstitutionSnies uint                       `gorm:"not null" mapper:"personHigherEducationInstitutionSnies"`
	HigherEducationInstitution      HigherEducationInstitution `gorm:"foreignKey:HigherEducationInstitutionSnies"`

	JobPosition string `gorm:"type:varchar(128); not null" mapper:"personJobPosition"`

	UserID    uint      `gorm:"not null" mapper:"userID"`
	User      User      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `mapper:"personCreatedAt"`
	UpdatedAt time.Time `mapper:"personUpdatedAt"`
	DeleteAt  time.Time `mapper:"personDeleteAt"`
}
