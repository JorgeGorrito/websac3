package usecase

import "websac3/app/domain/entity"

type GetReportFeedbackUseCase interface {
	Execute(reportID uint, lang string) (*entity.ReportFeedback, error)
}
