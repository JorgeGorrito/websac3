package request

type GetReportByIDRequest struct {
	ReportID uint `json:"report_id" validate:"required"`
}
