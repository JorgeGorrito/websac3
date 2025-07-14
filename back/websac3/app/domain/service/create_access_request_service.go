package service

import (
	"errors"
	"time"
	"websac3/app/domain/constants"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/domain/notification/template/context"
	"websac3/app/port/out/message"
	"websac3/app/port/out/notification"
	"websac3/app/port/out/notification/template"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/enum"
	"websac3/common/uuid"
)

type CreateAccessRequestService struct {
	msgProvider             message.Provider
	templateProvider        template.Provider
	createAccessRequestPort persistence.CreateAccessRequestPort
	updateAccessRequestPort persistence.UpdateAccessRequestPort
	createPersonPort        persistence.CreatePersonPort
	createUserPort          persistence.CreateUserPort
	getUserPort             persistence.GetUserPort
	updateUserPort          persistence.UpdateUserPort
	updatePersonPort        persistence.UpdatePersonPort
	getPersonPort           persistence.GetPersonPort
	getAccessRequestPort    persistence.GetAccessRequestPort
	statusEnum              enum.StatusEnum
	sendNotificationPort    notification.SendMailPort
	persistenceManager      persistence.Manager
}

func NewCreateAccessRequestService(
	createAccessRequestPort persistence.CreateAccessRequestPort,
	updateAccessRequestPort persistence.UpdateAccessRequestPort,
	createPersonPort persistence.CreatePersonPort,
	createUserPort persistence.CreateUserPort,
	updateUserPort persistence.UpdateUserPort,
	updatePersonPort persistence.UpdatePersonPort,
	getUserPort persistence.GetUserPort,
	getPersonPort persistence.GetPersonPort,
	getAccessRequestPort persistence.GetAccessRequestPort,
	statusEnum enum.StatusEnum,
	msgProvider message.Provider,
	notificationPort notification.SendMailPort,
	persistenceManager persistence.Manager,
	templateProvider template.Provider,
) *CreateAccessRequestService {
	return &CreateAccessRequestService{
		createAccessRequestPort: createAccessRequestPort,
		updateAccessRequestPort: updateAccessRequestPort,
		createPersonPort:        createPersonPort,
		createUserPort:          createUserPort,
		getUserPort:             getUserPort,
		getPersonPort:           getPersonPort,
		getAccessRequestPort:    getAccessRequestPort,
		statusEnum:              statusEnum,
		updateUserPort:          updateUserPort,
		updatePersonPort:        updatePersonPort,
		msgProvider:             msgProvider,
		sendNotificationPort:    notificationPort,
		persistenceManager:      persistenceManager,
		templateProvider:        templateProvider,
	}
}

func (c *CreateAccessRequestService) Execute(
	requestToCreate entity.AccessRequest,
	redirectURL string,
	lang string,
) error {
	return c.persistenceManager.ExecuteInTransaction(
		func(ctx persistence.Context) error {
			var err error
			var requestFound entity.AccessRequest
			var userFound entity.User
			var applicantToCreate *entity.Person = requestToCreate.Applicant
			var statusPending entity.Status

			if statusPending, err = c.statusEnum.GetByName(constants.Pending); err != nil {
				return err
			}
			requestToCreate.StatusID = statusPending.ID
			requestToCreate.Status = &statusPending

			if userFound, err = c.getUserPort.GetByDNI(
				applicantToCreate.IdentificationTypeID,
				applicantToCreate.IdentificationNumber,
				ctx,
			); err != nil {
				return err
			}

			if userFound.IsRegistered() {
				return errs.NewConflictError(
					c.msgProvider.
						WithLang(lang).
						GetMessage("create_access_request", "user_alredy_registered"),
				)
			}

			if requestFound, err = c.getAccessRequestPort.GetLastCreatedByIdentificationAndEmail(
				requestToCreate.Applicant.IdentificationTypeID,
				requestToCreate.Applicant.IdentificationNumber,
				requestToCreate.EmailValidation.To,
				ctx); err != nil {
				return err
			}

			if !requestFound.CanRegister() {
				return errs.NewConflictError(
					c.msgProvider.
						WithLang(lang).
						GetMessage("create_access_request", "cannot_create_access_request"),
				)
			}

			if !requestFound.IsApplicantRegistered() {
				if err = c.createPersonPort.Create(
					applicantToCreate,
					ctx,
				); err != nil {
					return err
				}
				requestToCreate.ApplicantID = applicantToCreate.ID
			} else {
				requestToCreate.ApplicantID = requestFound.ApplicantID
			}

			requestToCreate.ValidationCode = uuid.GenerateUUID()
			requestToCreate.IsVerified = false
			if err = c.createAccessRequestPort.Create(&requestToCreate, ctx); err != nil {
				return err
			}

			var templateEmail template.Template
			if templateEmail = c.templateProvider.GetByName("access_request_confirmation"); templateEmail == nil {
				return errors.New(
					c.msgProvider.
						WithLang(lang).
						GetMessage("mail_template", "template_not_found", "access_request_confirmation"),
				)
			}

			var templateToSend string
			if templateToSend, err = templateEmail.Render(
				&context.AccessRequestConfirmation{
					NombrePersona:      applicantToCreate.Name,
					EnlaceConfirmacion: redirectURL + requestToCreate.ValidationCode,
				},
			); err != nil {
				return err
			}

			requestToCreate.EmailValidation = &entity.EmailNotification{
				To:        requestToCreate.EmailValidation.To,
				Content:   templateToSend,
				Subject:   "Confirmación de solicitud de acceso",
				CreatedAt: time.Now(),
			}
			if err = c.sendNotificationPort.Send(
				requestToCreate.EmailValidation,
				ctx,
			); err != nil {
				return err
			}
			requestToCreate.EmailValidationID = func() *uint {
				if requestToCreate.EmailValidation == nil {
					return nil
				}
				id := requestToCreate.EmailValidation.ID
				return &id
			}()

			if err = c.updateAccessRequestPort.Update(&requestToCreate, ctx); err != nil {
				return err
			}

			return err
		},
	)
}
