package response

type GetCourseTopicsByCourseIDResponse struct {
	CourseID    uint                        `json:"course_id"`
	CourseName  string                      `json:"course_name"`
	CourseCode  string                      `json:"course_code"`
	Topics      []CourseTopicDetailResponse `json:"topics"`
	TotalTopics int                         `json:"total_topics"`
	TotalHours  float32                     `json:"total_hours"`
}

type CourseTopicDetailResponse struct {
	TopicID           uint    `json:"topic_id"`
	TopicName         string  `json:"topic_name"`
	StudyHours        float32 `json:"study_hours"`
	KnowledgeAreaID   uint    `json:"knowledge_area_id"`
	KnowledgeAreaName string  `json:"knowledge_area_name"`
}
