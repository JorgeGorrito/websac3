package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/paginator"
)

type ListReportFeedbacksService struct {
	listReportFeedbacksPort persistence.ListReportFeedbacksPort
	persistenceManager      db.Manager
	messageProvider         message.Provider
}

func NewListReportFeedbacksService(
	listReportFeedbacksPort persistence.ListReportFeedbacksPort,
	persistenceManager db.Manager,
	messageProvider message.Provider,
) *ListReportFeedbacksService {
	return &ListReportFeedbacksService{
		listReportFeedbacksPort: listReportFeedbacksPort,
		persistenceManager:      persistenceManager,
		messageProvider:         messageProvider,
	}
}

func (s *ListReportFeedbacksService) Execute(filters map[string]interface{}, page uint, limit uint, lang string) ([]entity.ReportFeedback, uint, error) {
	var result []entity.ReportFeedback
	var total uint
	err := s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		// Crear parámetros de paginación
		paginationParams := paginator.PaginationParams{
			Currentpage:  page,
			ItemsPerpage: limit,
		}

		// Obtener feedbacks con filtros y paginación
		feedbacks, totalCount, err := s.listReportFeedbacksPort.GetWithFilters(filters, paginationParams, ctx)
		if err != nil {
			return err
		}

		if len(feedbacks) == 0 {
			err = errs.NewNotFoundError(
				s.messageProvider.
					WithLang(lang).
					GetMessage("list_report_feedbacks", "not_found"),
			)
		}

		result = feedbacks
		total = totalCount
		return err
	})
	return result, total, err
}
