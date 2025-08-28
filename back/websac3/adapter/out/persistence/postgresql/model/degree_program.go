package model

type DegreeProgram struct {
	Snies               uint   `gorm:"primaryKey" json:"id"`
	Name                string `gorm:"not null; unique" json:"name"`
	TotalCredits        uint
	DurationValue       uint
	DurationUnit        DurationUnit
	ProgramFocus        string
	EntryProfile        string
	GraduateProfile     string
	ProfessionalProfile string
}
