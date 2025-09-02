package command

type CreateCourseTopicCommand struct {
	TopicID    uint `validations:"required"`
	StudyHours uint `validations:"required"`
}

type CreateCourseCommand struct {
	Name                        string                     `validations:"required"`
	Code                        string                     `validations:"required"`
	Credits                     uint                       `validations:"required"`
	PeriodNumber                uint                       `validations:"required"`
	NatureID                    uint                       `validations:"required"`
	TypeID                      uint                       `validations:"required"`
	IsCybersecurity             bool                       `validations:"required"`
	ContainsCybersecurityTopics bool                       `validations:"required"`
	DegreeProgramID             uint                       `validations:"required"`
	CourseTopics                []CreateCourseTopicCommand `validations:"optional"`
	CreatedBy                   uint                       `validations:"required"`
	Permissions                 []string
}
