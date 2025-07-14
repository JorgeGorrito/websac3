package request

type CreatePersonRequest struct {
	Name                            string `json:"name" `
	Lastname                        string `json:"lastname"  `
	IdentificationTypeID            uint   `json:"identification_type_id"  `
	IdentificationNumber            string `json:"identification_number"  `
	HigherEducationInstitutionSnies uint   `json:"higher_education_institution_snies"  `
	JobPosition                     string `json:"job_position"  `
	Email                           string `json:"email"  `
}
