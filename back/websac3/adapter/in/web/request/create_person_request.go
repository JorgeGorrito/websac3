package request

type CreatePersonRequest struct {
	Name                            string            `json:"name" binding:"required" mapper:"personName"`
	Lastname                        string            `json:"lastname" binding:"required" mapper:"personLastName"`
	IdentificationTypeID            uint              `json:"identification_type_id" binding:"required" mapper:"identificationTypeID"`
	IdentificationNumber            string            `json:"identification_number" binding:"required" mapper:"identificationNumber"`
	HigherEducationInstitutionSnies uint              `json:"higher_education_institution_snies" binding:"required" mapper:"higherEducationInstitutionSnies"`
	JobPosition                     string            `json:"job_position" binding:"required" mapper:"personJobPosition"`
	User                            CreateUserRequest `json:"user" mapper:"user" binding:"required"`
}
