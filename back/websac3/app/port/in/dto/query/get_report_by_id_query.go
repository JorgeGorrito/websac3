package query

type GetReportByIDQuery struct {
	ReportID uint `json:"report_id" validate:"required"`
}
