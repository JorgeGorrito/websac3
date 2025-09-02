package entity

type CourseTopic struct {
	CourseID   uint
	TopicID    uint
	StudyHours uint
}

type Course struct {
	ID                          uint
	Name                        string
	Code                        string
	Credits                     uint
	PeriodNumber                uint
	NatureID                    uint
	TypeID                      uint
	IsCybersecurity             bool
	ContainsCybersecurityTopics bool
	DegreeProgramID             uint
	DegreeProgram               *DegreeProgram
	CourseTopics                []CourseTopic
	CreatedBy                   uint
	UserCreator                 *User
}
