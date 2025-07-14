package command

type CreatePersonCommand struct {
	Name                            string `validations:"required"`
	Lastname                        string `validations:"required"`
	IdentificationNumber            string `validations:"required"`
	IdentificationTypeID            uint   `validations:"required"`
	HigherEducationInstitutionSnies uint   `validations:"required"`
	JobPosition                     string `validations:"required"`
	Email                           string `validations:"required;email"`
}
