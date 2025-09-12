package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type GetReportByIDService struct {
	getReportPort      persistence.GetReportPort
	msgProvider        message.Provider
	persistenceManager db.Manager
}

func NewGetReportByIDService(
	getReportPort persistence.GetReportPort,
	msgProvider message.Provider,
	persistenceManager db.Manager,
) *GetReportByIDService {
	return &GetReportByIDService{
		getReportPort:      getReportPort,
		msgProvider:        msgProvider,
		persistenceManager: persistenceManager,
	}
}

func (s *GetReportByIDService) Execute(reportID uint, lang string) (entity.Report, error) {
	var report entity.Report
	err := s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		// Obtener el reporte por ID
		var err error
		report, err = s.getReportPort.GetByID(reportID, ctx)
		if err != nil {
			return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("report", "not_found"))
		}

		return nil
	})

	if err != nil {
		return entity.Report{}, err
	}

	return report, nil
}
