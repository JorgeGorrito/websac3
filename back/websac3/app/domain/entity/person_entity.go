package entity

import "time"

type Person struct {
	ID       uint
	Name     string
	Lastname string

	IdentificationTypeID uint
	IdentificationType   *IdentificationType

	IdentificationNumber string

	HigherEducationInstitutionSnies uint
	HigherEducationInstitution      *HigherEducationInstitution

	JobPosition string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeleteAt    *time.Time
}

func (p *Person) IsRegistered() bool {
	return p.ID != 0
}
