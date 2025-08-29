package entity

type DegreeProgram struct {
	ID                  uint
	Snies               uint
	Name                string
	TotalCredits        uint
	DurationValue       uint
	DurationUnitID      uint
	DurationUnit        *DurationUnit
	ProgramFocus        string
	EntryProfile        string
	GraduateProfile     string
	ProfessionalProfile string
	CreatedBy           uint
	UserCreator         *User
}
