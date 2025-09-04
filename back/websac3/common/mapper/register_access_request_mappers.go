package mapper

import (
	"errors"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
)

func registerAccessRequestMappers() {
	RegisterMapFunc(func(createAccessRequestRequest *request.CreateAccessRequestRequest) (command.CreateAccessRequestCommand, error) {
		return command.CreateAccessRequestCommand{
			ValidationEmailURL:              createAccessRequestRequest.ValidationEmailURL,
			RegisterUserURL:                 createAccessRequestRequest.RegisterUserURL,
			Name:                            createAccessRequestRequest.Person.Name,
			Lastname:                        createAccessRequestRequest.Person.Lastname,
			IdentificationNumber:            createAccessRequestRequest.Person.IdentificationNumber,
			IdentificationTypeID:            createAccessRequestRequest.Person.IdentificationTypeID,
			HigherEducationInstitutionSnies: createAccessRequestRequest.Person.HigherEducationInstitutionSnies,
			JobPosition:                     createAccessRequestRequest.Person.JobPosition,
			Email:                           createAccessRequestRequest.Person.Email,
		}, nil
	})

	RegisterMapFunc(func(createAccessRequestCommand *command.CreateAccessRequestCommand) (entity.AccessRequest, error) {
		var errorList error
		return entity.AccessRequest{
			ValidationEmailURL: createAccessRequestCommand.ValidationEmailURL,
			CreateUserURL:      createAccessRequestCommand.RegisterUserURL,
			Applicant: &entity.Person{
				Name:                            createAccessRequestCommand.Name,
				Lastname:                        createAccessRequestCommand.Lastname,
				IdentificationNumber:            createAccessRequestCommand.IdentificationNumber,
				IdentificationTypeID:            createAccessRequestCommand.IdentificationTypeID,
				HigherEducationInstitutionSnies: createAccessRequestCommand.HigherEducationInstitutionSnies,
				JobPosition:                     createAccessRequestCommand.JobPosition,
			},
			EmailValidation: &entity.EmailNotification{
				To: createAccessRequestCommand.Email,
			},
		}, errorList
	})

	RegisterMapFunc(func(accessRequest *entity.AccessRequest) (model.AccessRequest, error) {
		return model.AccessRequest{
			ID:                       accessRequest.ID,
			ApplicantID:              accessRequest.ApplicantID,
			StatusID:                 accessRequest.StatusID,
			VerificationEmailID:      accessRequest.EmailValidationID,
			ApprovedEmailID:          accessRequest.EmailApprovedID,
			ApprovedRoleID:           accessRequest.ApprovedRoleID,
			ValidationEmailURL:       accessRequest.ValidationEmailURL,
			CreateUserURL:            accessRequest.CreateUserURL,
			ValidationEmailCode:      accessRequest.ValidationEmailCode,
			ValidationCreateUserCode: accessRequest.CreateUserCode,
			IsVerified:               accessRequest.IsVerified,
			CreatedAt:                accessRequest.CreatedAt,
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

		var approvedEmail *entity.EmailNotification
		var approvedEmailID *uint
		approvedEmailMapped, err := Map[model.Email, entity.EmailNotification](accessRequest.ApprovedEmail)
		if err != nil {
			if err != SrcPointerIsNilError {
				errorList = errors.Join(errorList, err)
			}
		} else {
			approvedEmail = &approvedEmailMapped
			approvedEmailID = &approvedEmailMapped.ID
		}

		var approvedRole *entity.Role
		var approvedRoleID *uint
		if accessRequest.ApprovedRoleID != nil {
			approvedRoleMapped, err := Map[model.Role, entity.Role](&accessRequest.ApprovedRole)
			if err != nil {
				errorList = errors.Join(errorList, err)
			} else {
				approvedRole = &approvedRoleMapped
				approvedRoleID = accessRequest.ApprovedRoleID
			}
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

			EmailApprovedID: approvedEmailID,
			EmailApproved:   approvedEmail,

			ApprovedRoleID: approvedRoleID,
			ApprovedRole:   approvedRole,

			ValidationEmailURL: accessRequest.ValidationEmailURL,
			CreateUserURL:      accessRequest.CreateUserURL,

			ValidationEmailCode: accessRequest.ValidationEmailCode,
			CreateUserCode:      accessRequest.ValidationCreateUserCode,

			IsVerified: accessRequest.IsVerified,
			CreatedAt:  accessRequest.CreatedAt,
		}, nil
	})

	RegisterMapFunc(func(accessRequest *entity.AccessRequest) (response.ListAccessRequestResponse, error) {
		var errorList error
		var listAccessRequestResponse response.ListAccessRequestResponse = response.ListAccessRequestResponse{
			ID:                                  accessRequest.ID,
			Name:                                accessRequest.Applicant.Name,
			Lastname:                            accessRequest.Applicant.Lastname,
			Email:                               accessRequest.EmailValidation.To,
			IdentificationNumber:                accessRequest.Applicant.IdentificationNumber,
			IdentificationType:                  accessRequest.Applicant.IdentificationType.Name,
			JobPosition:                         accessRequest.Applicant.JobPosition,
			HigherEducationInstitutionSnies:     accessRequest.Applicant.HigherEducationInstitutionSnies,
			HigherEducationInstitutionName:      accessRequest.Applicant.HigherEducationInstitution.Name,
			HigherEducationInstitutionOwnership: accessRequest.Applicant.HigherEducationInstitution.Ownership.Name,
			MunicipalityName:                    accessRequest.Applicant.HigherEducationInstitution.Municipality.Name,
			DepartmentName:                      accessRequest.Applicant.HigherEducationInstitution.Department.Name,
			StatusID:                            accessRequest.Status.ID,
			StatusName:                          accessRequest.Status.Name,
		}

		if errorList != nil {
			return response.ListAccessRequestResponse{}, errorList
		}

		return listAccessRequestResponse, nil
	})

	RegisterMapFunc(func(request *request.ApproveAccessRequestRequest) (command.ApproveAccessRequestCommand, error) {
		return command.ApproveAccessRequestCommand{
			RoleID: request.RoleID,
		}, nil
	})
}
