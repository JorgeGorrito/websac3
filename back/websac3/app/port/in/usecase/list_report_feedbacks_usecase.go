package usecase

import "websac3/app/domain/entity"

type ListReportFeedbacksUseCase interface {
	Execute(filters map[string]interface{}, page uint, limit uint, lang string) ([]entity.ReportFeedback, uint, error)
}
