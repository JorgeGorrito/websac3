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
			Lang:                     accessRequest.Lang,
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
			Lang:                accessRequest.Lang,
			IsVerified:          accessRequest.IsVerified,
			CreatedAt:           accessRequest.CreatedAt,
		}, nil
	})

	RegisterMapFunc(func(accessRequest *entity.AccessRequest) (response.ListAccessRequestResponse, error) {
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

		return listAccessRequestResponse, nil
	})

	RegisterMapFunc(func(request *request.ApproveAccessRequestRequest) (command.ApproveAccessRequestCommand, error) {
		return command.ApproveAccessRequestCommand{
			RoleID: request.RoleID,
		}, nil
	})

	// Entity to Response mapper for Approved Access Requests
	RegisterMapFunc(func(accessRequest *entity.AccessRequest) (response.ListApprovedAccessRequestResponse, error) {
		var applicantName, applicantLastname, applicantEmail, applicantIdentificationNumber, applicantIdentificationType, applicantJobPosition string
		var higherEducationInstitutionSnies uint
		var higherEducationInstitutionName, higherEducationInstitutionOwnership, municipalityName, departmentName string
		var statusID uint
		var statusName, approvedRoleName string
		var approvedRoleID uint
		var createdAt string

		if accessRequest.Applicant != nil {
			applicantName = accessRequest.Applicant.Name
			applicantLastname = accessRequest.Applicant.Lastname
			// Email no está disponible en Person, se obtiene del AccessRequest
			if accessRequest.EmailValidation != nil {
				applicantEmail = accessRequest.EmailValidation.To
			}
			applicantIdentificationNumber = accessRequest.Applicant.IdentificationNumber
			applicantJobPosition = accessRequest.Applicant.JobPosition

			if accessRequest.Applicant.IdentificationType != nil {
				applicantIdentificationType = accessRequest.Applicant.IdentificationType.Name
			}

			if accessRequest.Applicant.HigherEducationInstitution != nil {
				higherEducationInstitutionSnies = accessRequest.Applicant.HigherEducationInstitution.Snies
				higherEducationInstitutionName = accessRequest.Applicant.HigherEducationInstitution.Name

				if accessRequest.Applicant.HigherEducationInstitution.Ownership != nil {
					higherEducationInstitutionOwnership = accessRequest.Applicant.HigherEducationInstitution.Ownership.Name
				}

				if accessRequest.Applicant.HigherEducationInstitution.Municipality != nil {
					municipalityName = accessRequest.Applicant.HigherEducationInstitution.Municipality.Name
				}

				if accessRequest.Applicant.HigherEducationInstitution.Department != nil {
					departmentName = accessRequest.Applicant.HigherEducationInstitution.Department.Name
				}
			}
		}

		if accessRequest.Status != nil {
			statusID = accessRequest.Status.ID
			statusName = accessRequest.Status.Name
		}

		if accessRequest.ApprovedRole != nil {
			approvedRoleID = accessRequest.ApprovedRole.ID
			approvedRoleName = accessRequest.ApprovedRole.Name
		} else {
			// Si no hay rol aprobado, usar valores por defecto
			approvedRoleID = 0
			approvedRoleName = "Sin rol asignado"
		}

		createdAt = accessRequest.CreatedAt.Format("2006-01-02 15:04:05")

		return response.ListApprovedAccessRequestResponse{
			ID:                   accessRequest.ID,
			Name:                 applicantName,
			Email:                applicantEmail,
			Lastname:             applicantLastname,
			IdentificationNumber: applicantIdentificationNumber,
			IdentificationType:   applicantIdentificationType,
			JobPosition:          applicantJobPosition,

			HigherEducationInstitutionSnies:     higherEducationInstitutionSnies,
			HigherEducationInstitutionName:      higherEducationInstitutionName,
			HigherEducationInstitutionOwnership: higherEducationInstitutionOwnership,

			MunicipalityName: municipalityName,
			DepartmentName:   departmentName,
			StatusID:         statusID,
			StatusName:       statusName,
			ApprovedRoleID:   approvedRoleID,
			ApprovedRoleName: approvedRoleName,
			CreatedAt:        createdAt,
		}, nil
	})

	// Entity to Response mapper for Rejected Access Requests
	RegisterMapFunc(func(accessRequest *entity.AccessRequest) (response.ListRejectedAccessRequestResponse, error) {
		var applicantName, applicantLastname, applicantEmail, applicantIdentificationNumber, applicantIdentificationType, applicantJobPosition string
		var higherEducationInstitutionSnies uint
		var higherEducationInstitutionName, higherEducationInstitutionOwnership, municipalityName, departmentName string
		var statusID uint
		var statusName string
		var createdAt string

		if accessRequest.Applicant != nil {
			applicantName = accessRequest.Applicant.Name
			applicantLastname = accessRequest.Applicant.Lastname
			// Email no está disponible en Person, se obtiene del AccessRequest
			if accessRequest.EmailValidation != nil {
				applicantEmail = accessRequest.EmailValidation.To
			}
			applicantIdentificationNumber = accessRequest.Applicant.IdentificationNumber
			applicantJobPosition = accessRequest.Applicant.JobPosition

			if accessRequest.Applicant.IdentificationType != nil {
				applicantIdentificationType = accessRequest.Applicant.IdentificationType.Name
			}

			if accessRequest.Applicant.HigherEducationInstitution != nil {
				higherEducationInstitutionSnies = accessRequest.Applicant.HigherEducationInstitution.Snies
				higherEducationInstitutionName = accessRequest.Applicant.HigherEducationInstitution.Name

				if accessRequest.Applicant.HigherEducationInstitution.Ownership != nil {
					higherEducationInstitutionOwnership = accessRequest.Applicant.HigherEducationInstitution.Ownership.Name
				}

				if accessRequest.Applicant.HigherEducationInstitution.Municipality != nil {
					municipalityName = accessRequest.Applicant.HigherEducationInstitution.Municipality.Name
				}

				if accessRequest.Applicant.HigherEducationInstitution.Department != nil {
					departmentName = accessRequest.Applicant.HigherEducationInstitution.Department.Name
				}
			}
		}

		if accessRequest.Status != nil {
			statusID = accessRequest.Status.ID
			statusName = accessRequest.Status.Name
		}

		createdAt = accessRequest.CreatedAt.Format("2006-01-02 15:04:05")

		return response.ListRejectedAccessRequestResponse{
			ID:                   accessRequest.ID,
			Name:                 applicantName,
			Email:                applicantEmail,
			Lastname:             applicantLastname,
			IdentificationNumber: applicantIdentificationNumber,
			IdentificationType:   applicantIdentificationType,
			JobPosition:          applicantJobPosition,

			HigherEducationInstitutionSnies:     higherEducationInstitutionSnies,
			HigherEducationInstitutionName:      higherEducationInstitutionName,
			HigherEducationInstitutionOwnership: higherEducationInstitutionOwnership,

			MunicipalityName: municipalityName,
			DepartmentName:   departmentName,
			StatusID:         statusID,
			StatusName:       statusName,
			CreatedAt:        createdAt,
		}, nil
	})

	RegisterMapFunc(func(accessRequest *entity.AccessRequest) (response.UserAccessRequestResponse, error) {
		var errorList error
		var userAccessRequestResponse response.UserAccessRequestResponse = response.UserAccessRequestResponse{
			ID:                   accessRequest.ID,
			Name:                 accessRequest.Applicant.Name,
			Lastname:             accessRequest.Applicant.Lastname,
			IdentificationNumber: accessRequest.Applicant.IdentificationNumber,
			JobPosition:          accessRequest.Applicant.JobPosition,
			IsVerified:           accessRequest.IsVerified,
			CreatedAt:            accessRequest.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}

		// Mapear email
		if accessRequest.EmailValidation != nil {
			userAccessRequestResponse.Email = accessRequest.EmailValidation.To
		}

		// Mapear tipo de identificación
		if accessRequest.Applicant.IdentificationType != nil {
			userAccessRequestResponse.IdentificationType = accessRequest.Applicant.IdentificationType.Name
		}

		// Mapear institución de educación superior
		if accessRequest.Applicant.HigherEducationInstitution != nil {
			userAccessRequestResponse.HigherEducationInstitutionSnies = accessRequest.Applicant.HigherEducationInstitution.Snies
			userAccessRequestResponse.HigherEducationInstitutionName = accessRequest.Applicant.HigherEducationInstitution.Name

			// Mapear ownership
			if accessRequest.Applicant.HigherEducationInstitution.Ownership != nil {
				userAccessRequestResponse.HigherEducationInstitutionOwnership = accessRequest.Applicant.HigherEducationInstitution.Ownership.Name
			}

			// Mapear municipio y departamento
			if accessRequest.Applicant.HigherEducationInstitution.Municipality != nil {
				userAccessRequestResponse.MunicipalityName = accessRequest.Applicant.HigherEducationInstitution.Municipality.Name
			}

			// Mapear departamento directamente desde la institución
			if accessRequest.Applicant.HigherEducationInstitution.Department != nil {
				userAccessRequestResponse.DepartmentName = accessRequest.Applicant.HigherEducationInstitution.Department.Name
			}
		}

		// Mapear estado
		if accessRequest.Status != nil {
			userAccessRequestResponse.StatusID = accessRequest.Status.ID
			userAccessRequestResponse.StatusName = accessRequest.Status.Name
		}

		return userAccessRequestResponse, errorList
	})
}
