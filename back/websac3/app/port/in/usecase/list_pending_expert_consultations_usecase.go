package usecase

import (
	"websac3/app/domain/entity"
	commonfilter "websac3/common/filter"
	"websac3/common/paginator"
)

type ListPendingExpertConsultationsUseCase interface {
	Execute(paginationParams paginator.PaginationParams, filters commonfilter.Params, lang string) ([]entity.ExpertConsultation, uint, error)
}
