package service

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/paginator"
)

type ListReportFeedbacksService struct {
	listReportFeedbacksPort persistence.ListReportFeedbacksPort
	persistenceManager      db.Manager
}

func NewListReportFeedbacksService(
	listReportFeedbacksPort persistence.ListReportFeedbacksPort,
	persistenceManager db.Manager,
) *ListReportFeedbacksService {
	return &ListReportFeedbacksService{
		listReportFeedbacksPort: listReportFeedbacksPort,
		persistenceManager:      persistenceManager,
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

		result = feedbacks
		total = totalCount
		return nil
	})
	return result, total, err
}
