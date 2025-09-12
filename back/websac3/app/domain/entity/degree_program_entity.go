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

	Courses []Course
}

func (e *DegreeProgram) IsRegistered() bool {
	return e.ID != 0
}

func (e *DegreeProgram) GetCourseTopics() []CourseTopic {
	var courseTopics []CourseTopic = make([]CourseTopic, 0)
	for _, course := range e.Courses {
		courseTopics = append(courseTopics, course.GetCourseTopics()...)
	}
	return courseTopics
}
