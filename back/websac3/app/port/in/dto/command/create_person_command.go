package command

type CreatePersonCommand struct {
	Name                            string            `validate:"required" mapper:"personName"`
	Lastname                        string            `validate:"required" mapper:"personLastName"`
	IdentificationNumber            string            `validate:"required" mapper:"personIdentificationNumber"`
	IdentificationTypeID            uint              `validate:"required" mapper:"personIdentificationTypeID"`
	HigherEducationInstitutionSnies uint              `validate:"required" mapper:"personHigherEducationInstitutionSnies"`
	JobPosition                     string            `validate:"required" mapper:"personJobPosition"`
	User                            CreateUserCommand `mapper:"user"`
}
