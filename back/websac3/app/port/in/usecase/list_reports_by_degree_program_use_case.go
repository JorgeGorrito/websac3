package usecase

import "websac3/app/domain/entity"

type ListReportsByDegreeProgramUseCase interface {
	Execute(degreeProgramID uint, page, perPage uint, lang string) ([]entity.Report, int64, error)
}
