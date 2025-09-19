package entity

type CourseTopic struct {
	CourseID   uint
	TopicID    uint
	Topic      *Topic
	StudyHours float32
}

type Course struct {
	ID              uint
	Name            string
	Code            string
	Credits         uint
	PeriodNumber    uint
	NatureID        uint
	Nature          *CourseNature
	TypeID          uint
	Type            *CourseType
	IsCybersecurity bool

	DegreeProgramID uint
	DegreeProgram   *DegreeProgram

	CourseTopics []CourseTopic

	CreatedBy   uint
	UserCreator *User
}

func (e *Course) ContainsCybersecurityTopics() bool {
	return len(e.CourseTopics) > 0
}

func (e *Course) GetCourseTopics() []CourseTopic {
	return e.CourseTopics
}
