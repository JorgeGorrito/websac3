package response

type CourseTopicBulkResult struct {
	RowNumber  uint   `json:"row_number"`
	TopicID    uint   `json:"topic_id"`
	TopicName  string `json:"topic_name,omitempty"`
	StudyHours uint   `json:"study_hours"`
	Message    string `json:"message"`
}

type CourseTopicBulkError struct {
	RowNumber  uint     `json:"row_number"`
	TopicID    *uint    `json:"topic_id,omitempty"`
	StudyHours *uint    `json:"study_hours,omitempty"`
	Errors     []string `json:"errors"`
}

type BulkCreateCourseTopicResponse struct {
	TotalProcessed  uint                    `json:"total_processed"`
	SuccessfulCount uint                    `json:"successful_count"`
	FailedCount     uint                    `json:"failed_count"`
	SuccessfulItems []CourseTopicBulkResult `json:"successful_items"`
	FailedItems     []CourseTopicBulkError  `json:"failed_items"`
}

