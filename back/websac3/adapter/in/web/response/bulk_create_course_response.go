package response

type CourseBulkResult struct {
	RowNumber uint   `json:"row_number"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Message   string `json:"message"`
}

type CourseBulkError struct {
	RowNumber uint     `json:"row_number"`
	Code      *string  `json:"code,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Errors    []string `json:"errors"`
}

type BulkCreateCourseResponse struct {
	TotalProcessed  uint               `json:"total_processed"`
	SuccessfulCount uint               `json:"successful_count"`
	FailedCount     uint               `json:"failed_count"`
	SuccessfulItems []CourseBulkResult `json:"successful_items"`
	FailedItems     []CourseBulkError  `json:"failed_items"`
}







