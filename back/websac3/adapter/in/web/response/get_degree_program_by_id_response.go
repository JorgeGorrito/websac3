package response

// GetDegreeProgramByIDResponse representa la respuesta detallada de un programa de grado
type GetDegreeProgramByIDResponse struct {
	ID                  uint                           `json:"id" example:"1"`
	Snies               uint                           `json:"snies" example:"12345"`
	Name                string                         `json:"name" example:"Ingeniería de Sistemas"`
	TotalCredits        uint                           `json:"total_credits" example:"160"`
	DurationValue       uint                           `json:"duration_value" example:"5"`
	DurationUnit        ListDurationUnitResponse       `json:"duration_unit"`
	FormationLevel      ListFormationLevelsResponse    `json:"formation_level"`
	ProfessionalRoles   []ListProfessionalRoleResponse `json:"professional_roles"`
	ProgramFocus        string                         `json:"program_focus" example:"Formar ingenieros de sistemas con competencias en desarrollo de software"`
	EntryProfile        string                         `json:"entry_profile" example:"Bachiller con conocimientos básicos en matemáticas y lógica"`
	GraduateProfile     string                         `json:"graduate_profile" example:"Ingeniero de sistemas capaz de desarrollar soluciones tecnológicas"`
	ProfessionalProfile string                         `json:"professional_profile" example:"Desarrollador de software, analista de sistemas, arquitecto de software"`
	CreatedBy           uint                           `json:"created_by" example:"1"`
	CreatedAt           string                         `json:"created_at" example:"2024-01-15T10:30:00Z"`
	UpdatedAt           string                         `json:"updated_at" example:"2024-01-15T10:30:00Z"`
	// Información de la institución de educación superior
	HigherEducationInstitution *HigherEducationInstitutionInfo `json:"higher_education_institution,omitempty"`
}
