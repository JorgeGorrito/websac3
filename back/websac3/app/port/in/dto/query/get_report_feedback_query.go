package query

type GetReportFeedbackQuery struct {
	ReportID uint `json:"report_id" validate:"required"`

	Permissions []string `json:"permissions" validate:"required"`
}
