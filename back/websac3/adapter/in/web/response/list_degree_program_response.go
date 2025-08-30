package response

type ListDegreeProgramResponse struct {
	ID                  uint   `json:"id"`
	Snies               uint   `json:"snies"`
	Name                string `json:"name"`
	TotalCredits        uint   `json:"total_credits"`
	DurationValue       uint   `json:"duration_value"`
	DurationUnitID      uint   `json:"duration_unit_id"`
	ProgramFocus        string `json:"program_focus"`
	EntryProfile        string `json:"entry_profile"`
	GraduateProfile     string `json:"graduate_profile"`
	ProfessionalProfile string `json:"professional_profile"`
	CreatedBy           uint   `json:"created_by"`
}
