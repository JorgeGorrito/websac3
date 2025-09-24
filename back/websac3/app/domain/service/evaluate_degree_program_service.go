package service

import (
	"time"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/domain/notification/template/context"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/notification"
	"websac3/app/port/out/notification/template"
	"websac3/app/port/out/pdf"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
	"websac3/common/mail"
)

type EvaluateDegreeProgramService struct {
	getDegreeProgramPort    persistence.GetDegreeProgramPort
	getProfessionalRolePort persistence.GetProfessionalRolePort
	createReportPort        persistence.CreateReportPort
	getUserByIDUseCase      usecase.GetUserByIDUseCase
	msgProvider             message.Provider
	templateProvider        template.Provider
	pdfConverter            pdf.Converter
	sendNotificationPort    notification.SendMailPort
	persistenceManager      db.Manager
}

func NewEvaluateDegreeProgramService(
	getDegreeProgramPort persistence.GetDegreeProgramPort,
	getProfessionalRolePort persistence.GetProfessionalRolePort,
	createReportPort persistence.CreateReportPort,
	getUserByIDUseCase usecase.GetUserByIDUseCase,
	msgProvider message.Provider,
	templateProvider template.Provider,
	pdfConverter pdf.Converter,
	sendNotificationPort notification.SendMailPort,
	persistenceManager db.Manager,
) *EvaluateDegreeProgramService {
	return &EvaluateDegreeProgramService{
		getDegreeProgramPort:    getDegreeProgramPort,
		getProfessionalRolePort: getProfessionalRolePort,
		createReportPort:        createReportPort,
		getUserByIDUseCase:      getUserByIDUseCase,
		msgProvider:             msgProvider,
		templateProvider:        templateProvider,
		pdfConverter:            pdfConverter,
		sendNotificationPort:    sendNotificationPort,
		persistenceManager:      persistenceManager,
	}
}

func (s *EvaluateDegreeProgramService) Execute(degreeProgramID uint, professionalRoleID uint, userID uint, lang string) error {
	return s.persistenceManager.ExecuteInTransaction(func(ctx db.Context) error {
		// Obtener información del usuario para el correo
		user, err := s.getUserByIDUseCase.Execute(userID, lang)
		if err != nil {
			return err
		}

		if user.Email == "" {
			return errs.NewValidationError(s.msgProvider.WithLang(lang).GetMessage("evaluate_degree_program", "user_email_not_found"))
		}

		// Obtener programa de grado
		degreeProgram, err := s.getDegreeProgramPort.GetByIDWithLang(degreeProgramID, lang, ctx)
		if err != nil {
			return err
		}
		if !degreeProgram.IsRegistered() {
			return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("evaluate_degree_program", "not_found"))
		}

		// Obtener rol profesional
		professionalRole, err := s.getProfessionalRolePort.GetByID(professionalRoleID, lang, ctx)
		if err != nil {
			return err
		}
		if !professionalRole.IsRegistered() {
			return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("evaluate_degree_program", "professional_role_not_found"))
		}

		// Generar el reporte
		report := professionalRole.EvaluateDegreeProgram(&degreeProgram, lang)

		// Guardar el reporte en la base de datos
		savedReport, err := s.createReportPort.Create(report, ctx)
		if err != nil {
			return err
		}

		// Renderizar HTML del reporte
		reportTemplate := s.templateProvider.GetByNameAndLang("report", lang)
		if reportTemplate == nil {
			return errs.NewNotFoundError(s.msgProvider.WithLang(lang).GetMessage("mail_template", "template_not_found", "report"))
		}

		html, err := reportTemplate.Render(savedReport)
		if err != nil {
			return err
		}

		// Convertir a PDF
		pdfBytes, err := s.pdfConverter.HTMLToPDF(html)
		if err != nil {
			return err
		}

		// Crear adjunto PDF
		fileName := "reporte-evaluacion-" + time.Now().Format("20060102150405") + ".pdf"
		pdfAttachment := mail.CreateAttachmentFromBytes(pdfBytes, fileName, "application/pdf")

		// Preparar contenido del correo
		var emailContent string
		emailTemplate := s.templateProvider.GetByNameAndLang("degree_program_evaluation", lang)
		if emailTemplate == nil {
			// Usar template básico como fallback
			emailContent = `<html><body><h2>Hello, ` + user.Person.Name + `</h2><p>The degree program evaluation has been completed successfully.</p><p>Please find the detailed evaluation report attached to this email.</p><p>Best regards,<br>WebSAC3 System</p></body></html>`
		} else {
			// Renderizar contenido del correo usando la plantilla
			emailContent, err = emailTemplate.Render(&context.DegreeProgramEvaluation{
				NombrePersona: user.Person.Name,
			})
			if err != nil {
				return err
			}
		}

		// Crear notificación de email con adjunto
		emailNotification := &entity.EmailNotification{
			To:          user.Email,
			Subject:     s.msgProvider.WithLang(lang).GetMessage("evaluate_degree_program", "report_email_subject"),
			Content:     emailContent,
			CreatedAt:   time.Now(),
			Attachments: []entity.EmailAttachment{*pdfAttachment},
		}

		// Enviar correo
		return s.sendNotificationPort.Send(emailNotification, ctx)
	})
}
