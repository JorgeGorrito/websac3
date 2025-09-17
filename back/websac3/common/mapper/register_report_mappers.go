package mapper

import (
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
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
			// Mapear KnowledgeAreaReports sin relaciones complejas para evitar errores
			var knowledgeAreaReports []entity.KnowledgeAreaReport
			for _, kar := range reportModel.KnowledgeAreaReports {
				karEntity := entity.KnowledgeAreaReport{
					ID:                      kar.ID,
					Name:                    kar.Name,
					TotalLearnHoursExpected: kar.TotalLearnHoursExpected,
					TotalLearnHoursActual:   kar.TotalLearnHoursActual,
					ScoreExpected:           kar.ScoreExpected,
					ScoreGot:                kar.ScoreGot,
					// TopicReports se deja vacío para evitar errores de mapeo
				}
				knowledgeAreaReports = append(knowledgeAreaReports, karEntity)
			}

			// Mapear DegreeProgram sin relaciones complejas para evitar errores
			degreeProgram := entity.DegreeProgram{
				ID:                  reportModel.DegreeProgram.ID,
				Snies:               reportModel.DegreeProgram.Snies,
				Name:                reportModel.DegreeProgram.Name,
				TotalCredits:        reportModel.DegreeProgram.TotalCredits,
				DurationValue:       reportModel.DegreeProgram.DurationValue,
				DurationUnitID:      reportModel.DegreeProgram.DurationUnitID,
				ProgramFocus:        reportModel.DegreeProgram.ProgramFocus,
				EntryProfile:        reportModel.DegreeProgram.EntryProfile,
				GraduateProfile:     reportModel.DegreeProgram.GraduateProfile,
				ProfessionalProfile: reportModel.DegreeProgram.ProfessionalProfile,
				CreatedBy:           reportModel.DegreeProgram.CreatedBy,
				// Las relaciones complejas se dejan vacías para evitar errores de mapeo
			}

			// Mapear ProfessionalRole sin KnowledgeAreaExpected para evitar errores
			professionalRole := entity.ProfessionalRole{
				ID:   reportModel.ProfessionalRole.ID,
				Name: reportModel.ProfessionalRole.Name,
				// KnowledgeAreaExpected se deja vacío porque no se necesita para el reporte
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

	// KnowledgeAreaFeedback: model -> entity
	RegisterMapFunc(
		func(knowledgeAreaFeedbackModel *model.KnowledgeAreaFeedback) (entity.KnowledgeAreaFeedback, error) {
			// Mapear KnowledgeAreaReport si está disponible
			var knowledgeAreaReport *entity.KnowledgeAreaReport
			if knowledgeAreaFeedbackModel.KnowledgeAreaReport.ID != 0 {
				knowledgeAreaReport = &entity.KnowledgeAreaReport{
					ID:   knowledgeAreaFeedbackModel.KnowledgeAreaReport.ID,
					Name: knowledgeAreaFeedbackModel.KnowledgeAreaReport.Name,
				}
			}

			return entity.KnowledgeAreaFeedback{
				ID:                    knowledgeAreaFeedbackModel.ID,
				ReportFeedbackID:      knowledgeAreaFeedbackModel.ReportFeedbackID,
				KnowledgeAreaReportID: knowledgeAreaFeedbackModel.KnowledgeAreaReportID,
				KnowledgeAreaReport:   knowledgeAreaReport,
				Comments:              knowledgeAreaFeedbackModel.Comments,
				SecurityGaps:          knowledgeAreaFeedbackModel.SecurityGaps,
				Improvements:          knowledgeAreaFeedbackModel.Improvements,
				ComplianceLevel:       knowledgeAreaFeedbackModel.ComplianceLevel,
				AuditorRating:         knowledgeAreaFeedbackModel.AuditorRating,
			}, nil
		},
	)

	// KnowledgeAreaFeedback: entity -> model
	RegisterMapFunc(
		func(knowledgeAreaFeedbackEntity *entity.KnowledgeAreaFeedback) (model.KnowledgeAreaFeedback, error) {
			return model.KnowledgeAreaFeedback{
				ID:                    knowledgeAreaFeedbackEntity.ID,
				ReportFeedbackID:      knowledgeAreaFeedbackEntity.ReportFeedbackID,
				KnowledgeAreaReportID: knowledgeAreaFeedbackEntity.KnowledgeAreaReportID,
				Comments:              knowledgeAreaFeedbackEntity.Comments,
				SecurityGaps:          knowledgeAreaFeedbackEntity.SecurityGaps,
				Improvements:          knowledgeAreaFeedbackEntity.Improvements,
				ComplianceLevel:       knowledgeAreaFeedbackEntity.ComplianceLevel,
				AuditorRating:         knowledgeAreaFeedbackEntity.AuditorRating,
			}, nil
		},
	)

	// ReportFeedback: model -> entity
	RegisterMapFunc(
		func(reportFeedbackModel *model.ReportFeedback) (entity.ReportFeedback, error) {
			// Mapear KnowledgeAreaFeedbacks
			var knowledgeAreaFeedbacks []entity.KnowledgeAreaFeedback
			for _, kaf := range reportFeedbackModel.KnowledgeAreaFeedbacks {
				kafEntity, err := Map[model.KnowledgeAreaFeedback, entity.KnowledgeAreaFeedback](&kaf)
				if err != nil {
					return entity.ReportFeedback{}, err
				}
				knowledgeAreaFeedbacks = append(knowledgeAreaFeedbacks, kafEntity)
			}

			// Mapear Report (sin relaciones complejas para evitar errores)
			var report *entity.Report
			if reportFeedbackModel.Report.ID != 0 {
				report = &entity.Report{
					ID:    reportFeedbackModel.Report.ID,
					Score: reportFeedbackModel.Report.Score,
					DegreeProgram: entity.DegreeProgram{
						ID:    reportFeedbackModel.Report.DegreeProgram.ID,
						Name:  reportFeedbackModel.Report.DegreeProgram.Name,
						Snies: reportFeedbackModel.Report.DegreeProgram.Snies,
					},
					CreatedAt: reportFeedbackModel.Report.CreatedAt,
				}
			}

			// Mapear Auditor con Person
			var auditor *entity.User
			if reportFeedbackModel.Auditor.ID != 0 {
				auditor = &entity.User{
					ID:    reportFeedbackModel.Auditor.ID,
					Email: reportFeedbackModel.Auditor.Email,
				}
				// Mapear Person si está disponible
				if reportFeedbackModel.Auditor.Person.ID != 0 {
					auditor.Person = &entity.Person{
						ID:       reportFeedbackModel.Auditor.Person.ID,
						Name:     reportFeedbackModel.Auditor.Person.Name,
						Lastname: reportFeedbackModel.Auditor.Person.Lastname,
					}
				}
			}

			return entity.ReportFeedback{
				ID:                     reportFeedbackModel.ID,
				ReportID:               reportFeedbackModel.ReportID,
				Report:                 report,
				AuditorID:              reportFeedbackModel.AuditorID,
				Auditor:                auditor,
				GeneralComments:        reportFeedbackModel.GeneralComments,
				Recommendations:        reportFeedbackModel.Recommendations,
				AuditorRating:          reportFeedbackModel.AuditorRating,
				KnowledgeAreaFeedbacks: knowledgeAreaFeedbacks,
				CreatedAt:              reportFeedbackModel.CreatedAt,
				UpdatedAt:              reportFeedbackModel.UpdatedAt,
			}, nil
		},
	)

	// ReportFeedback: entity -> model
	RegisterMapFunc(
		func(reportFeedbackEntity *entity.ReportFeedback) (model.ReportFeedback, error) {
			// Mapear KnowledgeAreaFeedbacks
			var knowledgeAreaFeedbacks []model.KnowledgeAreaFeedback
			for _, kaf := range reportFeedbackEntity.KnowledgeAreaFeedbacks {
				kafModel, err := Map[entity.KnowledgeAreaFeedback, model.KnowledgeAreaFeedback](&kaf)
				if err != nil {
					return model.ReportFeedback{}, err
				}
				knowledgeAreaFeedbacks = append(knowledgeAreaFeedbacks, kafModel)
			}

			return model.ReportFeedback{
				ID:                     reportFeedbackEntity.ID,
				ReportID:               reportFeedbackEntity.ReportID,
				AuditorID:              reportFeedbackEntity.AuditorID,
				GeneralComments:        reportFeedbackEntity.GeneralComments,
				Recommendations:        reportFeedbackEntity.Recommendations,
				AuditorRating:          reportFeedbackEntity.AuditorRating,
				KnowledgeAreaFeedbacks: knowledgeAreaFeedbacks,
				CreatedAt:              reportFeedbackEntity.CreatedAt,
				UpdatedAt:              reportFeedbackEntity.UpdatedAt,
			}, nil
		},
	)

	// Request to Command mappers
	RegisterMapFunc(
		func(req *request.CreateReportFeedbackRequest) (command.CreateReportFeedbackCommand, error) {
			var knowledgeAreaFeedbacks []command.CreateKnowledgeAreaFeedbackCommand
			for _, kaf := range req.KnowledgeAreaFeedbacks {
				knowledgeAreaFeedbacks = append(knowledgeAreaFeedbacks, command.CreateKnowledgeAreaFeedbackCommand{
					KnowledgeAreaReportID: kaf.KnowledgeAreaReportID,
					Comments:              kaf.Comments,
					SecurityGaps:          kaf.SecurityGaps,
					Improvements:          kaf.Improvements,
					ComplianceLevel:       kaf.ComplianceLevel,
					AuditorRating:         kaf.AuditorRating,
				})
			}

			return command.CreateReportFeedbackCommand{
				ReportID:               req.ReportID,
				GeneralComments:        req.GeneralComments,
				Recommendations:        req.Recommendations,
				AuditorRating:          req.AuditorRating,
				KnowledgeAreaFeedbacks: knowledgeAreaFeedbacks,
				Permissions:            req.Permissions,
			}, nil
		},
	)

	// Command to Entity mappers
	RegisterMapFunc(
		func(cmd *command.CreateReportFeedbackCommand) (entity.ReportFeedback, error) {
			// Crear la entidad principal
			feedback := entity.NewReportFeedback(
				cmd.ReportID,
				cmd.UserID,
				cmd.GeneralComments,
				cmd.Recommendations,
				cmd.AuditorRating,
			)

			// Agregar knowledge area feedbacks
			for _, kaFeedbackCmd := range cmd.KnowledgeAreaFeedbacks {
				kaFeedback := entity.NewKnowledgeAreaFeedback(
					0, // Se asignará después de crear el feedback principal
					kaFeedbackCmd.KnowledgeAreaReportID,
					kaFeedbackCmd.Comments,
					kaFeedbackCmd.SecurityGaps,
					kaFeedbackCmd.Improvements,
					kaFeedbackCmd.ComplianceLevel,
					kaFeedbackCmd.AuditorRating,
				)
				feedback.AddKnowledgeAreaFeedback(kaFeedback)
			}

			return *feedback, nil
		},
	)

	// Request to Query mappers
	RegisterMapFunc(
		func(req *request.ListReportsPendingFeedbackRequest) (query.ListReportsPendingFeedbackQuery, error) {
			return query.ListReportsPendingFeedbackQuery{
				PaginationParams: req.PaginationParams,
				Permissions:      req.Permissions,
			}, nil
		},
	)

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

	// Entity to Response mappers for Report Feedback
	RegisterMapFunc(
		func(reportFeedbackEntity *entity.ReportFeedback) (response.ReportFeedbackResponse, error) {
			// Mapear KnowledgeAreaFeedbacks
			var knowledgeAreaFeedbacks []response.KnowledgeAreaFeedbackResponse
			for _, kaf := range reportFeedbackEntity.KnowledgeAreaFeedbacks {
				knowledgeAreaName := ""
				if kaf.KnowledgeAreaReport != nil {
					knowledgeAreaName = kaf.KnowledgeAreaReport.Name
				}

				knowledgeAreaFeedbacks = append(knowledgeAreaFeedbacks, response.KnowledgeAreaFeedbackResponse{
					ID:                    kaf.ID,
					KnowledgeAreaReportID: kaf.KnowledgeAreaReportID,
					KnowledgeAreaName:     knowledgeAreaName,
					Comments:              kaf.Comments,
					SecurityGaps:          kaf.SecurityGaps,
					Improvements:          kaf.Improvements,
					ComplianceLevel:       kaf.ComplianceLevel,
					AuditorRating:         kaf.AuditorRating,
				})
			}

			// Mapear Report Summary
			reportSummary := response.ReportSummaryResponse{
				ID:    reportFeedbackEntity.Report.ID,
				Score: reportFeedbackEntity.Report.Score,
				DegreeProgram: response.DegreeProgramSummaryResponse{
					ID:    reportFeedbackEntity.Report.DegreeProgram.ID,
					Name:  reportFeedbackEntity.Report.DegreeProgram.Name,
					Snies: reportFeedbackEntity.Report.DegreeProgram.Snies,
				},
				CreatedAt: reportFeedbackEntity.Report.CreatedAt.Format("2006-01-02T15:04:05Z"),
			}

			// Mapear Auditor Summary
			auditorName := ""
			if reportFeedbackEntity.Auditor.Person != nil {
				auditorName = reportFeedbackEntity.Auditor.Person.Name + " " + reportFeedbackEntity.Auditor.Person.Lastname
			}

			auditorSummary := response.UserSummaryResponse{
				ID:    reportFeedbackEntity.Auditor.ID,
				Email: reportFeedbackEntity.Auditor.Email,
				Name:  auditorName,
			}

			return response.ReportFeedbackResponse{
				ID:                     reportFeedbackEntity.ID,
				ReportID:               reportFeedbackEntity.ReportID,
				Report:                 reportSummary,
				Auditor:                auditorSummary,
				GeneralComments:        reportFeedbackEntity.GeneralComments,
				Recommendations:        reportFeedbackEntity.Recommendations,
				AuditorRating:          reportFeedbackEntity.AuditorRating,
				KnowledgeAreaFeedbacks: knowledgeAreaFeedbacks,
				CreatedAt:              reportFeedbackEntity.CreatedAt.Format("2006-01-02T15:04:05Z"),
				UpdatedAt:              reportFeedbackEntity.UpdatedAt.Format("2006-01-02T15:04:05Z"),
			}, nil
		},
	)

	// Report entity to ReportResponse (for pending reports list)
	RegisterMapFunc(
		func(reportEntity *entity.Report) (response.ReportResponse, error) {
			// Map DegreeProgram
			degreeProgramResp := response.DegreeProgramSummaryResponse{
				ID:    reportEntity.DegreeProgram.ID,
				Name:  reportEntity.DegreeProgram.Name,
				Snies: reportEntity.DegreeProgram.Snies,
			}

			// Map ProfessionalRole
			professionalRoleResp := response.ProfessionalRoleResponse{
				ID:   reportEntity.ProfessionalRole.ID,
				Name: reportEntity.ProfessionalRole.Name,
			}

			// Map HigherEducationInstitution (if available)
			var institutionResp response.InstitutionSummaryResponse
			if reportEntity.HigherEducationInstitution != nil && reportEntity.HigherEducationInstitution.Snies != 0 {
				institutionResp = response.InstitutionSummaryResponse{
					ID:   reportEntity.HigherEducationInstitution.Snies, // Usar Snies como ID
					Name: reportEntity.HigherEducationInstitution.Name,
				}
			}

			return response.ReportResponse{
				ID:                         reportEntity.ID,
				Score:                      reportEntity.Score,
				DegreeProgram:              degreeProgramResp,
				ProfessionalRole:           professionalRoleResp,
				HigherEducationInstitution: institutionResp,
				CreatedAt:                  reportEntity.CreatedAt.Format("2006-01-02T15:04:05Z"),
				HasFeedback:                false, // This will be set by the service based on business logic
			}, nil
		},
	)
}
