package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/paginator"
)

type ListReportsPendingFeedbackService struct {
	listReportsPendingFeedbackPort persistence.ListReportsPendingFeedbackPort
	getUserPort                    persistence.GetUserPort
	persistenceManager             db.Manager
	messageProvider                message.Provider
}

func NewListReportsPendingFeedbackService(
	listReportsPendingFeedbackPort persistence.ListReportsPendingFeedbackPort,
	getUserPort persistence.GetUserPort,
	persistenceManager db.Manager,
	messageProvider message.Provider,
) *ListReportsPendingFeedbackService {
	return &ListReportsPendingFeedbackService{
		listReportsPendingFeedbackPort: listReportsPendingFeedbackPort,
		getUserPort:                    getUserPort,
		persistenceManager:             persistenceManager,
		messageProvider:                messageProvider,
	}
}

func (s *ListReportsPendingFeedbackService) Execute(auditorID uint, paginationParams paginator.PaginationParams, lang string) ([]entity.Report, uint, error) {
	var reports []entity.Report = make([]entity.Report, 0)
	var total uint
	err := s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		// Obtener reportes que no tienen feedback
		var err error
		reports, total, err = s.listReportsPendingFeedbackPort.GetReportsWithoutFeedback(paginationParams, ctx)
		if err != nil {
			reports = nil
			return err
		}

		// Verificar si no se encontraron reportes
		if len(reports) == 0 {
			err = errs.NewNotFoundError(
				s.messageProvider.
					WithLang(lang).
					GetMessage("list_reports_pending_feedback", "not_found"),
			)
		}

		return err
	})
	return reports, total, err
}
