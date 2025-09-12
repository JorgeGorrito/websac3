package usecase

import "websac3/app/domain/entity"

type ListReportsByDegreeProgramUseCase interface {
	Execute(degreeProgramID uint, lang string) ([]entity.Report, error)
}
