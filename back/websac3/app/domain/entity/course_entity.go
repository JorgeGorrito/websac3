package entity

import "time"

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

	DeletedAt *time.Time
}

func (e *Course) ContainsCybersecurityTopics() bool {
	return len(e.CourseTopics) > 0
}

func (e *Course) GetCourseTopics() []CourseTopic {
	return e.CourseTopics
}

// CanBeDeletedBy verifica si un usuario puede eliminar este curso
// Un curso puede ser eliminado por:
// 1. El usuario que lo creó (CreatedBy)
// 2. Un administrador (verificado a través de User.IsAdmin())
func (e *Course) CanBeDeletedBy(userID uint) bool {
	return e.CreatedBy == userID
}
