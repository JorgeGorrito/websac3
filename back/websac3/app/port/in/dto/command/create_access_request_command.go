package command

type CreateAccessRequestCommand struct {
	Name                            string   `validations:"required"`
	Lastname                        string   `validations:"required"`
	IdentificationNumber            string   `validations:"required"`
	IdentificationTypeID            uint     `validations:"required"`
	HigherEducationInstitutionSnies uint     `validations:"required"`
	JobPosition                     string   `validations:"required"`
	Email                           string   `validations:"required;email"`
	ValidationEmailURL              string   `validations:"required"`
	RegisterUserURL                 string   `validations:"required"`
	Permissions                     []string
}
