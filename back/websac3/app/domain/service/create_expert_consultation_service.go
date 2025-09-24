package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type CreateExpertConsultationService struct {
	createExpertConsultationPort persistence.CreateExpertConsultationPort
	getUserPort                  persistence.GetUserPort
	getDegreeProgramPort         persistence.GetDegreeProgramPort
	getReportPort                persistence.GetReportPort
	msgProvider                  message.Provider
	persistenceManager           db.Manager
}

func NewCreateExpertConsultationService(
	createExpertConsultationPort persistence.CreateExpertConsultationPort,
	getUserPort persistence.GetUserPort,
	getDegreeProgramPort persistence.GetDegreeProgramPort,
	getReportPort persistence.GetReportPort,
	msgProvider message.Provider,
	persistenceManager db.Manager,
) *CreateExpertConsultationService {
	return &CreateExpertConsultationService{
		createExpertConsultationPort: createExpertConsultationPort,
		getUserPort:                  getUserPort,
		getDegreeProgramPort:         getDegreeProgramPort,
		getReportPort:                getReportPort,
		msgProvider:                  msgProvider,
		persistenceManager:           persistenceManager,
	}
}

func (c *CreateExpertConsultationService) Execute(
	consultationToCreate entity.ExpertConsultation,
	lang string,
) error {
	return c.persistenceManager.ExecuteInTransaction(
		func(ctx db.Context) error {
			var err error
			var userFound entity.User
			var degreeProgramFound entity.DegreeProgram
			var reportFound entity.Report

			// Validar que el usuario solicitante existe
			if userFound, err = c.getUserPort.GetByID(consultationToCreate.RequesterID, ctx); err != nil {
				return err
			}

			if !userFound.IsRegistered() {
				return errs.NewNotFoundError(
					c.msgProvider.
						WithLang(lang).
						GetMessage("create_expert_consultation", "user_not_found"),
				)
			}

			// Validar que el programa de grado existe
			if degreeProgramFound, err = c.getDegreeProgramPort.GetByID(consultationToCreate.DegreeProgramID, ctx); err != nil {
				return err
			}

			if !degreeProgramFound.IsRegistered() {
				return errs.NewNotFoundError(
					c.msgProvider.
						WithLang(lang).
						GetMessage("create_expert_consultation", "degree_program_not_found"),
				)
			}

			// Validar que el reporte existe
			if reportFound, err = c.getReportPort.GetByID(consultationToCreate.ReportID, ctx); err != nil {
				return err
			}

			if reportFound.ID == 0 {
				return errs.NewNotFoundError(
					c.msgProvider.
						WithLang(lang).
						GetMessage("create_expert_consultation", "report_not_found"),
				)
			}

			// Establecer el estado inicial como pendiente
			consultationToCreate.StatusID = 1 // ID for "pending" status

			// Crear la consulta de experto
			if err = c.createExpertConsultationPort.Create(&consultationToCreate, ctx); err != nil {
				return err
			}

			return nil
		},
	)
}
