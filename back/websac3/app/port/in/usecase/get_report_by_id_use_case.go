package usecase

import "websac3/app/domain/entity"

type GetReportByIDUseCase interface {
	Execute(reportID uint, lang string) (entity.Report, error)
}
