package command

type CreateDegreeProgramCommand struct {
	Snies               uint   `validations:"required"`
	Name                string `validations:"required"`
	TotalCredits        uint   `validations:"required"`
	DurationValue       uint   `validations:"required"`
	DurationUnitID      uint   `validations:"required"`
	FormationLevelID    uint   `validations:"required"`
	ProfessionalRoleIDs []uint `validations:"required"`
	ProgramFocus        string `validations:"required"`
	EntryProfile        string `validations:"required"`
	GraduateProfile     string `validations:"required"`
	ProfessionalProfile string `validations:"required"`
	CreatedBy           uint   `validations:"required"`
	Permissions         []string
}
