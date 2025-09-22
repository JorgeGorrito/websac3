package command

type UpdateCourseTopicCommand struct {
	TopicID    uint `validations:"required"`
	StudyHours uint `validations:"required"`
}

type UpdateCourseCommand struct {
	ID                          uint                       `validations:"required"`
	Name                        string                     `validations:"required"`
	Code                        string                     `validations:"required"`
	Credits                     uint                       `validations:"required"`
	PeriodNumber                uint                       `validations:"required"`
	NatureID                    uint                       `validations:"required"`
	TypeID                      uint                       `validations:"required"`
	IsCybersecurity             bool                       `validations:"required"`
	ContainsCybersecurityTopics bool                       `validations:"required"`
	DegreeProgramID             uint                       `validations:"required"`
	CourseTopics                []UpdateCourseTopicCommand `validations:"optional"`
	UserID                      uint                       `validations:"required"`
	Permissions                 []string
}
