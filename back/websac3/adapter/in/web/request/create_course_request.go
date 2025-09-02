package request

type CreateCourseTopicRequest struct {
	TopicID    uint `json:"topic_id" example:"1"`
	StudyHours uint `json:"study_hours" example:"20"`
}

type CreateCourseRequest struct {
	Name                        string                     `json:"name" example:"Programación Avanzada"`
	Code                        string                     `json:"code" example:"PROG301"`
	Credits                     uint                       `json:"credits" example:"3"`
	PeriodNumber                uint                       `json:"period_number" example:"3"`
	NatureID                    uint                       `json:"nature_id" example:"1"`
	TypeID                      uint                       `json:"type_id" example:"1"`
	IsCybersecurity             bool                       `json:"is_cybersecurity" example:"false"`
	ContainsCybersecurityTopics bool                       `json:"contains_cybersecurity_topics" example:"true"`
	DegreeProgramID             uint                       `json:"degree_program_id" example:"1"`
	CourseTopics                []CreateCourseTopicRequest `json:"course_topics,omitempty"`
}
