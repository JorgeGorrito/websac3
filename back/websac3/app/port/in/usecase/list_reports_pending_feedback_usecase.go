package usecase

import (
	"websac3/app/domain/entity"
	"websac3/common/paginator"
)

type ListReportsPendingFeedbackUseCase interface {
	Execute(auditorID uint, paginationParams paginator.PaginationParams, lang string) ([]entity.Report, uint, error)
}
