package mapper

import (
	"errors"
	"time"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
)

func registerExpertConsultationMappers() {
	// Registrar mapper de ExpertConsultationStatus primero
	RegisterMapFunc(func(expertConsultationStatus *model.ExpertConsultationStatus) (entity.ExpertConsultationStatus, error) {
		var name string
		var names []entity.ExpertConsultationStatusName

		if len(expertConsultationStatus.Names) > 0 {
			name = expertConsultationStatus.Names[0].Name
			// Mapear todos los nombres
			for _, statusName := range expertConsultationStatus.Names {
				names = append(names, entity.ExpertConsultationStatusName{
					ID:   statusName.ID,
					Lang: statusName.Lang,
					Name: statusName.Name,
				})
			}
		}

		return entity.ExpertConsultationStatus{
			ID:    expertConsultationStatus.ID,
			Name:  name,
			Names: names,
		}, nil
	})

	RegisterMapFunc(func(createExpertConsultationRequest *request.CreateExpertConsultationRequest) (command.CreateExpertConsultationCommand, error) {
		return command.CreateExpertConsultationCommand{
			// RequesterID se asigna en el controlador desde el token JWT
			DegreeProgramID: createExpertConsultationRequest.DegreeProgramID,
			ReportID:        createExpertConsultationRequest.ReportID,
			RequestMessage:  createExpertConsultationRequest.RequestMessage,
		}, nil
	})

	RegisterMapFunc(func(createExpertConsultationCommand *command.CreateExpertConsultationCommand) (entity.ExpertConsultation, error) {
		return entity.ExpertConsultation{
			RequesterID:     createExpertConsultationCommand.RequesterID,
			DegreeProgramID: createExpertConsultationCommand.DegreeProgramID,
			ReportID:        createExpertConsultationCommand.ReportID,
			RequestMessage:  createExpertConsultationCommand.RequestMessage,
		}, nil
	})

	RegisterMapFunc(func(expertConsultation *entity.ExpertConsultation) (model.ExpertConsultation, error) {
		var updatedAt time.Time
		if expertConsultation.UpdatedAt != nil {
			updatedAt = *expertConsultation.UpdatedAt
		}

		return model.ExpertConsultation{
			ID:              expertConsultation.ID,
			RequesterID:     expertConsultation.RequesterID,
			DegreeProgramID: expertConsultation.DegreeProgramID,
			ReportID:        expertConsultation.ReportID,
			RequestMessage:  expertConsultation.RequestMessage,
			ExpertResponse:  expertConsultation.ExpertResponse,
			ExpertID:        expertConsultation.ExpertID,
			StatusID:        expertConsultation.StatusID,
			CreatedAt:       expertConsultation.CreatedAt,
			UpdatedAt:       updatedAt,
			AnsweredAt:      expertConsultation.AnsweredAt,
			ClosedAt:        expertConsultation.ClosedAt,
		}, nil
	})

	RegisterMapFunc(func(expertConsultation *model.ExpertConsultation) (entity.ExpertConsultation, error) {
		var errorList error

		// Mapear Requester
		var requester *entity.User
		if expertConsultation.Requester.ID != 0 {
			requesterMapped, err := Map[model.User, entity.User](&expertConsultation.Requester)
			if err != nil {
				errorList = errors.Join(errorList, err)
			} else {
				requester = &requesterMapped
			}
		}

		// Mapear DegreeProgram
		var degreeProgram *entity.DegreeProgram
		if expertConsultation.DegreeProgram.ID != 0 {
			degreeProgramMapped, err := Map[model.DegreeProgram, entity.DegreeProgram](&expertConsultation.DegreeProgram)
			if err != nil {
				errorList = errors.Join(errorList, err)
			} else {
				degreeProgram = &degreeProgramMapped
			}
		}

		// Mapear Report
		var report *entity.Report
		if expertConsultation.Report.ID != 0 {
			reportMapped, err := Map[model.Report, entity.Report](&expertConsultation.Report)
			if err != nil {
				errorList = errors.Join(errorList, err)
			} else {
				report = &reportMapped
			}
		}

		// Mapear Expert
		var expert *entity.User
		if expertConsultation.Expert != nil && expertConsultation.Expert.ID != 0 {
			expertMapped, err := Map[model.User, entity.User](expertConsultation.Expert)
			if err != nil {
				errorList = errors.Join(errorList, err)
			} else {
				expert = &expertMapped
			}
		}

		// Mapear Status
		var status *entity.ExpertConsultationStatus
		if expertConsultation.Status.ID != 0 {
			statusMapped, err := Map[model.ExpertConsultationStatus, entity.ExpertConsultationStatus](&expertConsultation.Status)
			if err != nil {
				errorList = errors.Join(errorList, err)
			} else {
				status = &statusMapped
			}
		}

		if errorList != nil {
			return entity.ExpertConsultation{}, errorList
		}

		return entity.ExpertConsultation{
			ID:              expertConsultation.ID,
			RequesterID:     expertConsultation.RequesterID,
			Requester:       requester,
			DegreeProgramID: expertConsultation.DegreeProgramID,
			DegreeProgram:   degreeProgram,
			ReportID:        expertConsultation.ReportID,
			Report:          report,
			RequestMessage:  expertConsultation.RequestMessage,
			ExpertResponse:  expertConsultation.ExpertResponse,
			ExpertID:        expertConsultation.ExpertID,
			Expert:          expert,
			StatusID:        expertConsultation.StatusID,
			Status:          status,
			CreatedAt:       expertConsultation.CreatedAt,
			UpdatedAt:       &expertConsultation.UpdatedAt,
			AnsweredAt:      expertConsultation.AnsweredAt,
			ClosedAt:        expertConsultation.ClosedAt,
		}, nil
	})

	RegisterMapFunc(func(expertConsultation *entity.ExpertConsultation) (response.UserExpertConsultationResponse, error) {
		var userExpertConsultationResponse response.UserExpertConsultationResponse = response.UserExpertConsultationResponse{
			ID:          expertConsultation.ID,
			RequesterID: expertConsultation.RequesterID,
			CreatedAt:   expertConsultation.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}

		// Mapear requester
		if expertConsultation.Requester != nil {
			userExpertConsultationResponse.RequesterName = expertConsultation.Requester.Email
			userExpertConsultationResponse.RequesterEmail = expertConsultation.Requester.Email
			if expertConsultation.Requester.Person != nil {
				userExpertConsultationResponse.RequesterName = expertConsultation.Requester.Person.Name + " " + expertConsultation.Requester.Person.Lastname
			}
		}

		// Mapear degree program
		if expertConsultation.DegreeProgram != nil {
			userExpertConsultationResponse.DegreeProgramID = expertConsultation.DegreeProgram.ID
			userExpertConsultationResponse.DegreeProgramName = expertConsultation.DegreeProgram.Name
			userExpertConsultationResponse.DegreeProgramSnies = expertConsultation.DegreeProgram.Snies
		}

		// Mapear report
		if expertConsultation.Report != nil {
			userExpertConsultationResponse.ReportID = expertConsultation.Report.ID
			userExpertConsultationResponse.ReportScore = expertConsultation.Report.Score
		}

		// Mapear mensajes
		userExpertConsultationResponse.RequestMessage = expertConsultation.RequestMessage
		userExpertConsultationResponse.ExpertResponse = expertConsultation.ExpertResponse

		// Mapear expert
		if expertConsultation.Expert != nil {
			userExpertConsultationResponse.ExpertID = expertConsultation.ExpertID
			userExpertConsultationResponse.ExpertEmail = &expertConsultation.Expert.Email
			if expertConsultation.Expert.Person != nil {
				expertName := expertConsultation.Expert.Person.Name + " " + expertConsultation.Expert.Person.Lastname
				userExpertConsultationResponse.ExpertName = &expertName
			}
		}

		// Mapear status
		if expertConsultation.Status != nil {
			userExpertConsultationResponse.StatusID = expertConsultation.Status.ID
			userExpertConsultationResponse.StatusName = expertConsultation.Status.Name
		}

		// Mapear fechas
		if expertConsultation.UpdatedAt != nil {
			updatedAt := expertConsultation.UpdatedAt.Format("2006-01-02T15:04:05Z")
			userExpertConsultationResponse.UpdatedAt = &updatedAt
		}
		if expertConsultation.AnsweredAt != nil {
			answeredAt := expertConsultation.AnsweredAt.Format("2006-01-02T15:04:05Z")
			userExpertConsultationResponse.AnsweredAt = &answeredAt
		}
		if expertConsultation.ClosedAt != nil {
			closedAt := expertConsultation.ClosedAt.Format("2006-01-02T15:04:05Z")
			userExpertConsultationResponse.ClosedAt = &closedAt
		}

		return userExpertConsultationResponse, nil
	})

	RegisterMapFunc(func(expertConsultation *entity.ExpertConsultation) (response.PendingExpertConsultationResponse, error) {
		var pendingExpertConsultationResponse response.PendingExpertConsultationResponse = response.PendingExpertConsultationResponse{
			ID:          expertConsultation.ID,
			RequesterID: expertConsultation.RequesterID,
			CreatedAt:   expertConsultation.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}

		// Mapear requester
		if expertConsultation.Requester != nil {
			pendingExpertConsultationResponse.RequesterEmail = expertConsultation.Requester.Email
			if expertConsultation.Requester.Person != nil {
				pendingExpertConsultationResponse.RequesterName = expertConsultation.Requester.Person.Name + " " + expertConsultation.Requester.Person.Lastname
				pendingExpertConsultationResponse.RequesterJobPosition = expertConsultation.Requester.Person.JobPosition

				// Mapear información de la institución educativa
				if expertConsultation.Requester.Person.HigherEducationInstitution != nil {
					pendingExpertConsultationResponse.RequesterInstitutionSnies = expertConsultation.Requester.Person.HigherEducationInstitution.Snies
					pendingExpertConsultationResponse.RequesterInstitutionName = expertConsultation.Requester.Person.HigherEducationInstitution.Name

					// Mapear ownership de la institución
					if expertConsultation.Requester.Person.HigherEducationInstitution.Ownership != nil {
						pendingExpertConsultationResponse.RequesterInstitutionOwnership = expertConsultation.Requester.Person.HigherEducationInstitution.Ownership.Name
					}
				}
			}
		}

		// Mapear degree program
		if expertConsultation.DegreeProgram != nil {
			pendingExpertConsultationResponse.DegreeProgramID = expertConsultation.DegreeProgram.ID
			pendingExpertConsultationResponse.DegreeProgramName = expertConsultation.DegreeProgram.Name
			pendingExpertConsultationResponse.DegreeProgramSnies = expertConsultation.DegreeProgram.Snies
		}

		// Mapear report
		if expertConsultation.Report != nil {
			pendingExpertConsultationResponse.ReportID = expertConsultation.Report.ID
			pendingExpertConsultationResponse.ReportScore = expertConsultation.Report.Score
		}

		// Mapear mensaje de solicitud
		pendingExpertConsultationResponse.RequestMessage = expertConsultation.RequestMessage

		// Mapear status
		if expertConsultation.Status != nil {
			pendingExpertConsultationResponse.StatusID = expertConsultation.Status.ID
			pendingExpertConsultationResponse.StatusName = expertConsultation.Status.Name
		}

		// Mapear fecha de actualización
		if expertConsultation.UpdatedAt != nil {
			updatedAt := expertConsultation.UpdatedAt.Format("2006-01-02T15:04:05Z")
			pendingExpertConsultationResponse.UpdatedAt = &updatedAt
		}

		return pendingExpertConsultationResponse, nil
	})

	// Mappers para aceptar y rechazar consultas
	RegisterMapFunc(func(acceptRequest *request.AcceptExpertConsultationRequest) (command.AcceptExpertConsultationCommand, error) {
		return command.AcceptExpertConsultationCommand{
			ExpertResponse: acceptRequest.ExpertResponse,
		}, nil
	})

	RegisterMapFunc(func(rejectRequest *request.RejectExpertConsultationRequest) (command.RejectExpertConsultationCommand, error) {
		return command.RejectExpertConsultationCommand{
			ExpertResponse: rejectRequest.ExpertResponse,
		}, nil
	})
}
