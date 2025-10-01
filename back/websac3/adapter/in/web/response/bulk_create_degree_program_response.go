package response

type BulkCreateDegreeProgramResponse struct {
	TotalProcessed  int                       `json:"total_processed"`
	SuccessfulCount int                       `json:"successful_count"`
	FailedCount     int                       `json:"failed_count"`
	SuccessfulItems []DegreeProgramBulkResult `json:"successful_items"`
	FailedItems     []DegreeProgramBulkError  `json:"failed_items"`
}

type DegreeProgramBulkResult struct {
	RowNumber int    `json:"row_number"`
	Snies     uint   `json:"snies"`
	Name      string `json:"name"`
	Message   string `json:"message"`
}

type DegreeProgramBulkError struct {
	RowNumber int      `json:"row_number"`
	Snies     *uint    `json:"snies,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Errors    []string `json:"errors"`
}
