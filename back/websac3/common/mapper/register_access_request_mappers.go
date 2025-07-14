package mapper

import (
	"errors"
	"websac3/adapter/in/web/request"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
)

func registerAccessRequestMappers() {
	RegisterMapFunc(func(createAccessRequestRequest *request.CreateAccessRequestRequest) (command.CreateAccessRequestCommand, error) {
		var errorList error

		personMapped, err := Map[request.CreatePersonRequest, command.CreatePersonCommand](&createAccessRequestRequest.Person)
		if err != nil {
			errorList = errors.Join(errorList, err)
		}

		return command.CreateAccessRequestCommand{
			RedirectUrlTo: createAccessRequestRequest.RedirectUrlTo,
			Person:        personMapped,
		}, errorList
	})

	RegisterMapFunc(func(createAccessRequestCommand *command.CreateAccessRequestCommand) (entity.AccessRequest, error) {
		var errorList error

		personMapped, err := Map[command.CreatePersonCommand, entity.Person](&createAccessRequestCommand.Person)
		if err != nil {
			errorList = errors.Join(errorList, err)
		}

		return entity.AccessRequest{
			Applicant: &personMapped,
			EmailValidation: &entity.EmailNotification{
				To: createAccessRequestCommand.Person.Email,
			},
		}, errorList
	})

	RegisterMapFunc(func(accessRequest *entity.AccessRequest) (model.AccessRequest, error) {
		return model.AccessRequest{
			ID:                  accessRequest.ID,
			ApplicantID:         accessRequest.ApplicantID,
			StatusID:            accessRequest.StatusID,
			VerificationEmailID: accessRequest.EmailValidationID,
			ValidationCode:      accessRequest.ValidationCode,
			IsVerified:          accessRequest.IsVerified,
			CreatedAt:           accessRequest.CreatedAt,
		}, nil
	})

	RegisterMapFunc(func(accessRequest *model.AccessRequest) (entity.AccessRequest, error) {
		var errorList error
		applicantMapped, err := Map[model.Person, entity.Person](&accessRequest.Applicant)
		if err != nil {
			errorList = errors.Join(errorList, err)
		}
		statusMapped, err := Map[model.Status, entity.Status](&accessRequest.Status)
		if err != nil {
			errorList = errors.Join(errorList, err)
		}

		var verificationEmail *entity.EmailNotification
		var verificationEmailID *uint
		verificationEmailMapped, err := Map[model.Email, entity.EmailNotification](accessRequest.VerificationEmail)
		if err != nil {
			if err != SrcPointerIsNilError {
				errorList = errors.Join(errorList, err)
			}
		} else {
			verificationEmail = &verificationEmailMapped
			verificationEmailID = &verificationEmailMapped.ID
		}

		if errorList != nil {
			return entity.AccessRequest{}, errorList
		}

		return entity.AccessRequest{
			ID: accessRequest.ID,

			ApplicantID: applicantMapped.ID,
			Applicant:   &applicantMapped,

			StatusID: accessRequest.StatusID,
			Status:   &statusMapped,

			EmailValidationID: verificationEmailID,
			EmailValidation:   verificationEmail,

			ValidationCode: accessRequest.ValidationCode,
			IsVerified:     accessRequest.IsVerified,
			CreatedAt:      accessRequest.CreatedAt,
		}, nil
	})
}
