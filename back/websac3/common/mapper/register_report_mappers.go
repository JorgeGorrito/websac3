package mapper

import (
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/query"
)

func registerReportMappers() {
	// TopicReport: entity -> model
	RegisterMapFunc(
		func(topicReportEntity *entity.TopicReport) (model.TopicReport, error) {
			return model.TopicReport{
				ID:                 topicReportEntity.ID,
				TopicID:            topicReportEntity.TopicID,
				Name:               topicReportEntity.Name,
				LearnHoursExpected: topicReportEntity.LearnHoursExpected,
				LearnHoursActual:   topicReportEntity.LearnHoursActual,
			}, nil
		},
	)

	// TopicReport: model -> entity
	RegisterMapFunc(
		func(topicReportModel *model.TopicReport) (entity.TopicReport, error) {
			return entity.TopicReport{
				ID:                 topicReportModel.ID,
				TopicID:            topicReportModel.TopicID,
				Name:               topicReportModel.Name,
				LearnHoursExpected: topicReportModel.LearnHoursExpected,
				LearnHoursActual:   topicReportModel.LearnHoursActual,
			}, nil
		},
	)

	// KnowledgeAreaReport: entity -> model
	RegisterMapFunc(
		func(knowledgeAreaReportEntity *entity.KnowledgeAreaReport) (model.KnowledgeAreaReport, error) {
			var topicReports []model.TopicReport
			for _, tr := range knowledgeAreaReportEntity.TopicReports {
				trModel, err := Map[entity.TopicReport, model.TopicReport](&tr)
				if err != nil {
					return model.KnowledgeAreaReport{}, err
				}
				topicReports = append(topicReports, trModel)
			}

			return model.KnowledgeAreaReport{
				ID:                      knowledgeAreaReportEntity.ID,
				Name:                    knowledgeAreaReportEntity.Name,
				TotalLearnHoursExpected: knowledgeAreaReportEntity.TotalLearnHoursExpected,
				TotalLearnHoursActual:   knowledgeAreaReportEntity.TotalLearnHoursActual,
				ScoreExpected:           knowledgeAreaReportEntity.ScoreExpected,
				ScoreGot:                knowledgeAreaReportEntity.ScoreGot,
				TopicReports:            topicReports,
			}, nil
		},
	)

	// KnowledgeAreaReport: model -> entity
	RegisterMapFunc(
		func(knowledgeAreaReportModel *model.KnowledgeAreaReport) (entity.KnowledgeAreaReport, error) {
			var topicReports []entity.TopicReport
			for _, tr := range knowledgeAreaReportModel.TopicReports {
				trEntity, err := Map[model.TopicReport, entity.TopicReport](&tr)
				if err != nil {
					return entity.KnowledgeAreaReport{}, err
				}
				topicReports = append(topicReports, trEntity)
			}

			return entity.KnowledgeAreaReport{
				ID:                      knowledgeAreaReportModel.ID,
				Name:                    knowledgeAreaReportModel.Name,
				TotalLearnHoursExpected: knowledgeAreaReportModel.TotalLearnHoursExpected,
				TotalLearnHoursActual:   knowledgeAreaReportModel.TotalLearnHoursActual,
				ScoreExpected:           knowledgeAreaReportModel.ScoreExpected,
				ScoreGot:                knowledgeAreaReportModel.ScoreGot,
				TopicReports:            topicReports,
			}, nil
		},
	)

	// Report: entity -> model
	RegisterMapFunc(
		func(reportEntity *entity.Report) (model.Report, error) {
			var knowledgeAreaReports []model.KnowledgeAreaReport
			for _, kar := range reportEntity.KnowledgeAreaReports {
				karModel, err := Map[entity.KnowledgeAreaReport, model.KnowledgeAreaReport](&kar)
				if err != nil {
					return model.Report{}, err
				}
				knowledgeAreaReports = append(knowledgeAreaReports, karModel)
			}

			return model.Report{
				ID:                   reportEntity.ID,
				ProfessionalRoleID:   reportEntity.ProfessionalRole.ID,
				DegreeProgramID:      reportEntity.DegreeProgram.ID,
				Score:                reportEntity.Score,
				KnowledgeAreaReports: knowledgeAreaReports,
			}, nil
		},
	)

	// Report: model -> entity
	RegisterMapFunc(
		func(reportModel *model.Report) (entity.Report, error) {
			var knowledgeAreaReports []entity.KnowledgeAreaReport
			for _, kar := range reportModel.KnowledgeAreaReports {
				karEntity, err := Map[model.KnowledgeAreaReport, entity.KnowledgeAreaReport](&kar)
				if err != nil {
					return entity.Report{}, err
				}
				knowledgeAreaReports = append(knowledgeAreaReports, karEntity)
			}

			// Map DegreeProgram and ProfessionalRole
			degreeProgram, err := Map[model.DegreeProgram, entity.DegreeProgram](&reportModel.DegreeProgram)
			if err != nil {
				return entity.Report{}, err
			}

			professionalRole, err := Map[model.ProfessionalRole, entity.ProfessionalRole](&reportModel.ProfessionalRole)
			if err != nil {
				return entity.Report{}, err
			}

			return entity.Report{
				ID:                   reportModel.ID,
				DegreeProgram:        degreeProgram,
				ProfessionalRole:     professionalRole,
				KnowledgeAreaReports: knowledgeAreaReports,
				Score:                reportModel.Score,
				CreatedAt:            reportModel.CreatedAt,
				// Note: HigherEducationInstitution would need to be reconstructed from related data if needed
			}, nil
		},
	)

	// Request to Query mappers
	RegisterMapFunc(
		func(req *request.ListReportsByDegreeProgramRequest) (query.ListReportsByDegreeProgramQuery, error) {
			return query.ListReportsByDegreeProgramQuery{
				DegreeProgramID: req.DegreeProgramID,
			}, nil
		},
	)

	RegisterMapFunc(
		func(req *request.GetReportByIDRequest) (query.GetReportByIDQuery, error) {
			return query.GetReportByIDQuery{
				ReportID: req.ReportID,
			}, nil
		},
	)

	// Entity to Response mappers
	RegisterMapFunc(
		func(reportEntity *entity.Report) (response.ListReportResponse, error) {
			// Map DegreeProgram
			degreeProgramResp, err := Map[entity.DegreeProgram, response.ListDegreeProgramResponse](&reportEntity.DegreeProgram)
			if err != nil {
				return response.ListReportResponse{}, err
			}

			// Map ProfessionalRole
			professionalRoleResp := response.ProfessionalRoleResponse{
				ID:   reportEntity.ProfessionalRole.ID,
				Name: reportEntity.ProfessionalRole.Name,
			}

			// Map KnowledgeAreaReports
			var knowledgeAreaReports []response.KnowledgeAreaReportResponse
			for _, kar := range reportEntity.KnowledgeAreaReports {
				var topicReports []response.TopicReportResponse
				for _, tr := range kar.TopicReports {
					topicReports = append(topicReports, response.TopicReportResponse{
						ID:                 tr.ID,
						TopicID:            tr.TopicID,
						Name:               tr.Name,
						LearnHoursExpected: tr.LearnHoursExpected,
						LearnHoursActual:   tr.LearnHoursActual,
					})
				}

				knowledgeAreaReports = append(knowledgeAreaReports, response.KnowledgeAreaReportResponse{
					ID:                      kar.ID,
					Name:                    kar.Name,
					TotalLearnHoursExpected: kar.TotalLearnHoursExpected,
					TotalLearnHoursActual:   kar.TotalLearnHoursActual,
					ScoreExpected:           kar.ScoreExpected,
					ScoreGot:                kar.ScoreGot,
					TopicReports:            topicReports,
				})
			}

			return response.ListReportResponse{
				ID:                   reportEntity.ID,
				DegreeProgramID:      reportEntity.DegreeProgram.ID,
				DegreeProgram:        degreeProgramResp,
				ProfessionalRole:     professionalRoleResp,
				Score:                reportEntity.Score,
				CreatedAt:            reportEntity.CreatedAt,
				KnowledgeAreaReports: knowledgeAreaReports,
			}, nil
		},
	)
}
