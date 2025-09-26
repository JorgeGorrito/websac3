package request

type CreateDegreeProgramRequest struct {
	Snies               uint   `json:"snies" example:"12345"`
	Name                string `json:"name" example:"Ingeniería de Sistemas"`
	TotalCredits        uint   `json:"total_credits" example:"160"`
	DurationValue       uint   `json:"duration_value" example:"5"`
	DurationUnitID      uint   `json:"duration_unit_id" example:"1"`
	FormationLevelID    uint   `json:"formation_level_id" example:"3"`
	ProfessionalRoleIDs []uint `json:"professional_role_ids"`
	ProgramFocus        string `json:"program_focus" example:"Desarrollo de software"`
	EntryProfile        string `json:"entry_profile" example:"Bachiller con conocimientos básicos en matemáticas"`
	GraduateProfile     string `json:"graduate_profile" example:"Ingeniero capaz de desarrollar sistemas informáticos"`
	ProfessionalProfile string `json:"professional_profile" example:"Desarrollador de software, analista de sistemas"`
}
