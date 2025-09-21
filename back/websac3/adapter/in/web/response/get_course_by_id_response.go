package response

type GetCourseByIDResponse struct {
	ID                          uint                  `json:"id"`
	Name                        string                `json:"name"`
	Code                        string                `json:"code"`
	Credits                     uint                  `json:"credits"`
	PeriodNumber                uint                  `json:"period_number"`
	NatureID                    uint                  `json:"nature_id"`
	NatureName                  string                `json:"nature_name"`
	TypeID                      uint                  `json:"type_id"`
	TypeName                    string                `json:"type_name"`
	IsCybersecurity             bool                  `json:"is_cybersecurity"`
	ContainsCybersecurityTopics bool                  `json:"contains_cybersecurity_topics"`
	DegreeProgramID             uint                  `json:"degree_program_id"`
	DegreeProgramName           string                `json:"degree_program_name"`
	CreatedBy                   uint                  `json:"created_by"`
	CreatorName                 string                `json:"creator_name"`
	CourseTopics                []CourseTopicResponse `json:"course_topics"`
}

type CourseTopicResponse struct {
	TopicID    uint    `json:"topic_id"`
	TopicName  string  `json:"topic_name"`
	StudyHours float32 `json:"study_hours"`
}
