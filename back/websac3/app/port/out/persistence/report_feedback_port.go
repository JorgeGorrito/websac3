package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
	"websac3/common/paginator"
)

type CreateReportFeedbackPort interface {
	Create(feedback *entity.ReportFeedback, ctx db.Context) error
}

type GetReportFeedbackPort interface {
	GetByID(feedbackID uint, ctx db.Context) (*entity.ReportFeedback, error)
	GetByReportID(reportID uint, ctx db.Context) (*entity.ReportFeedback, error)
}

type UpdateReportFeedbackPort interface {
	Update(feedback *entity.ReportFeedback, ctx db.Context) error
}

type ListReportFeedbacksPort interface {
	GetWithFilters(filters map[string]interface{}, paginationParams paginator.PaginationParams, ctx db.Context) ([]entity.ReportFeedback, uint, error)
}

type ListReportsPendingFeedbackPort interface {
	GetReportsWithoutFeedback(paginationParams paginator.PaginationParams, ctx db.Context) ([]entity.Report, uint, error)
}
