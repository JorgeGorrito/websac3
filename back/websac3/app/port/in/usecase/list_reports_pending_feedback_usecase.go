package usecase

import (
	"websac3/app/domain/entity"
	"websac3/common/filter"
	"websac3/common/paginator"
)

type ListReportsPendingFeedbackUseCase interface {
	Execute(auditorID uint, paginationParams paginator.PaginationParams, filters filter.Params, sortBy, sortOrder, lang string) ([]entity.Report, uint, error)
}
