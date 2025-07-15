package request

type CreatePersonRequest struct {
	Name                            string `json:"name" example:"Jorge"`
	Lastname                        string `json:"lastname" example:"Gorrito"`
	IdentificationTypeID            uint   `json:"identification_type_id" example:"1"`
	IdentificationNumber            string `json:"identification_number" example:"123456789"`
	HigherEducationInstitutionSnies uint   `json:"higher_education_institution_snies" example:"1119"`
	JobPosition                     string `json:"job_position"  example:"director de programa"`
	Email                           string `json:"email" example:"j0rg3.4b3ll4@gmail.com"`
}
