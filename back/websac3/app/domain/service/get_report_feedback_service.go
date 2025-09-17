package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type GetReportFeedbackService struct {
	getReportFeedbackPort persistence.GetReportFeedbackPort
	persistenceManager    db.Manager
}

func NewGetReportFeedbackService(
	getReportFeedbackPort persistence.GetReportFeedbackPort,
	persistenceManager db.Manager,
) *GetReportFeedbackService {
	return &GetReportFeedbackService{
		getReportFeedbackPort: getReportFeedbackPort,
		persistenceManager:    persistenceManager,
	}
}

func (s *GetReportFeedbackService) Execute(reportID uint, lang string) (*entity.ReportFeedback, error) {
	var result *entity.ReportFeedback
	err := s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		// Obtener feedback por ID de reporte
		feedback, err := s.getReportFeedbackPort.GetByReportID(reportID, ctx)
		if err != nil {
			return errs.NewNotFoundError("feedback_not_found")
		}
		result = feedback
		return nil
	})
	return result, err
}
