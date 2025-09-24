package mapper

import (
	"errors"
	"time"
	"websac3/adapter/in/web/request"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
)

func registerExpertConsultationMappers() {
	RegisterMapFunc(func(createExpertConsultationRequest *request.CreateExpertConsultationRequest) (command.CreateExpertConsultationCommand, error) {
		return command.CreateExpertConsultationCommand{
			RequesterID:     createExpertConsultationRequest.RequesterID,
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
}
