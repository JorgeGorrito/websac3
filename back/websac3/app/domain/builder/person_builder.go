package builder

import "websac3/app/domain/entity"

type PersonBuilder interface {
	WithID(id uint) PersonBuilder
	WithName(name string) PersonBuilder
	WithLastname(lastname string) PersonBuilder
	WithIdentificationTypeID(identificationTypeID uint) PersonBuilder
	WithIdentificationNumber(identificationNumber string) PersonBuilder
	WithHigherEducationInstitutionSnies(higherEducationInstitutionSnies uint) PersonBuilder
	WithJobPosition(jobPosition string) PersonBuilder
	WithUser(user *entity.User) PersonBuilder
	Build() *entity.Person
}

type personBuilder struct {
	person entity.Person
}

func NewPersonBuilder() PersonBuilder {
	return &personBuilder{}
}

func (b *personBuilder) WithID(id uint) PersonBuilder {
	b.person.ID = id
	return b
}

func (b *personBuilder) WithName(name string) PersonBuilder {
	b.person.Name = name
	return b
}

func (b *personBuilder) WithLastname(lastname string) PersonBuilder {
	b.person.Lastname = lastname
	return b
}

func (b *personBuilder) WithIdentificationTypeID(identificationTypeID uint) PersonBuilder {
	b.person.IdentificationTypeID = identificationTypeID
	return b
}

func (b *personBuilder) WithIdentificationNumber(identificationNumber string) PersonBuilder {
	b.person.IdentificationNumber = identificationNumber
	return b
}

func (b *personBuilder) WithHigherEducationInstitutionSnies(higherEducationInstitutionSnies uint) PersonBuilder {
	b.person.HigherEducationInstitutionSnies = higherEducationInstitutionSnies
	return b
}

func (b *personBuilder) WithJobPosition(jobPosition string) PersonBuilder {
	b.person.JobPosition = jobPosition
	return b
}

func (b *personBuilder) WithUser(user *entity.User) PersonBuilder {
	b.person.User = user
	return b
}

func (b *personBuilder) Build() *entity.Person {
	return &b.person
}
