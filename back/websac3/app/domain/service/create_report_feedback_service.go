package service

import (
	"fmt"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/domain/notification/template/context"
	"websac3/app/port/out/message"
	"websac3/app/port/out/notification"
	"websac3/app/port/out/notification/template"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type CreateReportFeedbackService struct {
	getReportPort            persistence.GetReportPort
	getUserPort              persistence.GetUserPort
	createReportFeedbackPort persistence.CreateReportFeedbackPort
	getReportFeedbackPort    persistence.GetReportFeedbackPort
	sendNotificationPort     notification.SendMailPort
	persistenceManager       db.Manager
	msgProvider              message.Provider
	templateProvider         template.Provider
}

func NewCreateReportFeedbackService(
	getReportPort persistence.GetReportPort,
	getUserPort persistence.GetUserPort,
	createReportFeedbackPort persistence.CreateReportFeedbackPort,
	getReportFeedbackPort persistence.GetReportFeedbackPort,
	sendNotificationPort notification.SendMailPort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
	templateProvider template.Provider,
) *CreateReportFeedbackService {
	return &CreateReportFeedbackService{
		getReportPort:            getReportPort,
		getUserPort:              getUserPort,
		createReportFeedbackPort: createReportFeedbackPort,
		getReportFeedbackPort:    getReportFeedbackPort,
		sendNotificationPort:     sendNotificationPort,
		persistenceManager:       persistenceManager,
		msgProvider:              msgProvider,
		templateProvider:         templateProvider,
	}
}

func (s *CreateReportFeedbackService) sendFeedbackNotification(report *entity.Report, feedback *entity.ReportFeedback, auditor *entity.User, lang string, ctx db.Context) {

	// Crear contexto para el template
	templateContext := &context.ReportFeedbackContext{
		ProgramLeadName:   report.DegreeProgram.UserCreator.Person.Name + " " + report.DegreeProgram.UserCreator.Person.Lastname,
		AuditorName:       auditor.Person.Name + " " + auditor.Person.Lastname,
		DegreeProgramName: report.DegreeProgram.Name,
		ReportScore:       report.Score,
		GeneralComments:   feedback.GeneralComments,
		Recommendations:   feedback.Recommendations,
		InstitutionName:   report.HigherEducationInstitution.Name,
	}

	// Renderizar template
	templateName := "report_feedback"

	template := s.templateProvider.GetByNameAndLang(templateName, lang)
	if template == nil {
		// Intentar con el template sin idioma
		template = s.templateProvider.GetByName(templateName)
		if template == nil {
			panic(fmt.Sprintf("Template '%s' no está registrado", templateName))
		}
	}

	htmlContent, err := template.Render(templateContext)
	if err != nil {
		panic(fmt.Sprintf("Error al renderizar template: %v", err))
	}

	// Crear notificación de email
	emailNotification := &entity.EmailNotification{
		To:          report.DegreeProgram.UserCreator.Email,
		Subject:     s.msgProvider.WithLang(lang).GetMessage("report_feedback", "subject"),
		Content:     htmlContent,
		Attachments: []entity.EmailAttachment{}, // Sin adjuntos por ahora
	}

	// Enviar notificación
	s.sendNotificationPort.Send(emailNotification, ctx)
}

func (s *CreateReportFeedbackService) Execute(reportID uint, auditorID uint, feedback *entity.ReportFeedback, lang string) error {
	return s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		// Verificar que el reporte existe
		report, err := s.getReportPort.GetByID(reportID, ctx)
		if err != nil {
			return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("report_feedback", "report_not_found"))
		}

		// Verificar que el usuario es cybersecurity auditor
		auditor, err := s.getUserPort.GetByID(auditorID, ctx)
		if err != nil {
			return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("report_feedback", "auditor_not_found"))
		}

		// Verificar que no existe feedback previo para este reporte
		existingFeedback, _ := s.getReportFeedbackPort.GetByReportID(reportID, ctx)
		if existingFeedback != nil {
			return errs.NewConflictError(s.msgProvider.WithLang(lang).GetMessage("report_feedback", "feedback_already_exists"))
		}

		// Crear el feedback
		feedback.ReportID = reportID
		feedback.AuditorID = auditorID

		err = s.createReportFeedbackPort.Create(feedback, ctx)
		if err != nil {
			return err
		}

		// Notificar al creator del programa sobre el feedback
		if report.DegreeProgram.UserCreator != nil && report.DegreeProgram.UserCreator.Email != "" {
			s.sendFeedbackNotification(&report, feedback, &auditor, lang, ctx)
		}

		return nil
	})
}
