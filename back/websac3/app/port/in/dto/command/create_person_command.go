package command

import "websac3/common/validator"

type CreatePersonCommand struct {
	Name                            string            `validate:"required" mapper:"personName"`
	Lastname                        string            `validate:"required" mapper:"personLastName"`
	IdentificationNumber            string            `validate:"required" mapper:"identificationNumber"`
	IdentificationTypeID            uint              `validate:"required" mapper:"identificationTypeID"`
	HigherEducationInstitutionSnies uint              `validate:"required" mapper:"higherEducationInstitutionSnies"`
	JobPosition                     string            `validate:"required" mapper:"personJobPosition"`
	User                            CreateUserCommand `mapper:"user"`
}

func (p *CreatePersonCommand) Validate() error {
	return validator.ValidateFields(p)
}
