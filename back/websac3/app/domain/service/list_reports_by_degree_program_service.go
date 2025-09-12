package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type ListReportsByDegreeProgramService struct {
	getReportPort        persistence.GetReportPort
	getDegreeProgramPort persistence.GetDegreeProgramPort
	msgProvider          message.Provider
	persistenceManager   db.Manager
}

func NewListReportsByDegreeProgramService(
	getReportPort persistence.GetReportPort,
	getDegreeProgramPort persistence.GetDegreeProgramPort,
	msgProvider message.Provider,
	persistenceManager db.Manager,
) *ListReportsByDegreeProgramService {
	return &ListReportsByDegreeProgramService{
		getReportPort:        getReportPort,
		getDegreeProgramPort: getDegreeProgramPort,
		msgProvider:          msgProvider,
		persistenceManager:   persistenceManager,
	}
}

func (s *ListReportsByDegreeProgramService) Execute(degreeProgramID uint, lang string) ([]entity.Report, error) {
	var reports []entity.Report
	err := s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		// Verificar que el programa de grado existe
		degreeProgram, err := s.getDegreeProgramPort.GetByID(degreeProgramID, ctx)
		if err != nil {
			return err
		}
		if !degreeProgram.IsRegistered() {
			return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("degree_program", "not_found"))
		}

		// Obtener los reportes del programa de grado
		reports, err = s.getReportPort.GetByDegreeProgramID(degreeProgramID, ctx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return reports, nil
}
