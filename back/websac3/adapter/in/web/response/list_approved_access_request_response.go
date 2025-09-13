package response

type ListApprovedAccessRequestResponse struct {
	ID                   uint   `json:"id"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Lastname             string `json:"lastname"`
	IdentificationNumber string `json:"identification_number"`
	IdentificationType   string `json:"identification_type"`
	JobPosition          string `json:"job_position"`

	HigherEducationInstitutionSnies     uint   `json:"higher_education_institution_snies"`
	HigherEducationInstitutionName      string `json:"higher_education_institution_name"`
	HigherEducationInstitutionOwnership string `json:"higher_education_institution_ownership"`

	MunicipalityName string `json:"municipality_name"`
	DepartmentName   string `json:"department_name"`
	StatusID         uint   `json:"status_id"`
	StatusName       string `json:"status_name"`
	ApprovedRoleID   uint   `json:"approved_role_id"`
	ApprovedRoleName string `json:"approved_role_name"`
	CreatedAt        string `json:"created_at"`
}
