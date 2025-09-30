package response

type GetUserProfileResponse struct {
	ID                     uint   `json:"id"`
	Email                  string `json:"email"`
	Name                   string `json:"name"`
	Lastname               string `json:"lastname"`
	IdentificationTypeID   uint   `json:"identification_type_id"`
	IdentificationTypeName string `json:"identification_type_name"`
	IdentificationNumber   string `json:"identification_number"`
	InstitutionSnies       uint   `json:"institution_snies"`
	InstitutionName        string `json:"institution_name"`
	JobPosition            string `json:"job_position"`
	RoleID                 uint   `json:"role_id"`
	RoleName               string `json:"role_name"`
	IsActive               bool   `json:"is_active"`
}
