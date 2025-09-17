package usecase

import "websac3/app/domain/entity"

type CreateReportFeedbackUseCase interface {
	Execute(reportID uint, auditorID uint, feedback *entity.ReportFeedback, lang string) error
}
