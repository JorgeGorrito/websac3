package usecase

import "websac3/app/domain/entity"

type UpdateReportFeedbackUseCase interface {
	Execute(feedbackID uint, feedback *entity.ReportFeedback, lang string) error
}
