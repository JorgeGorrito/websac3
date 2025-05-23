package entity

import "time"

type Person struct {
	ID                              uint      `mapper:"personID"`
	Name                            string    `mapper:"personName"`
	Lastname                        string    `mapper:"personLastname"`
	IdentificationTypeID            uint      `mapper:"personIdentificationTypeID"`
	IdentificationNumber            string    `mapper:"personIdentificationNumber"`
	HigherEducationInstitutionSnies uint      `mapper:"personHigherEducationInstitutionSnies"`
	JobPosition                     string    `mapper:"personJobPosition"`
	User                            *User     `mapper:"user"`
	CreatedAt                       time.Time `mapper:"personCreatedAt"`
	UpdatedAt                       time.Time `mapper:"personUpdatedAt"`
	DeleteAt                        time.Time `mapper:"personDeleteAt"`
}

func (p *Person) IsRegistered() bool {
	return p.ID != 0
}
