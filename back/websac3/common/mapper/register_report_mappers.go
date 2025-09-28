package mapper

import (
	"math"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/dto/query"
)

// roundScore rounds a float32 score to 4 decimal places
func roundScore(score float32) float32 {
	return float32(math.Round(float64(score)*10000) / 10000)
}

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
				ScoreExpected:           roundScore(knowledgeAreaReportEntity.ScoreExpected),
				ScoreGot:                roundScore(knowledgeAreaReportEntity.ScoreGot),
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
				ScoreExpected:           roundScore(knowledgeAreaReportModel.ScoreExpected),
				ScoreGot:                roundScore(knowledgeAreaReportModel.ScoreGot),
				TopicReports:            topicReports,
			}, nil
		},
	)

	// UnexpectedTopicReport: entity -> model
	RegisterMapFunc(
		func(unexpectedTopicReportEntity *entity.UnexpectedTopicReport) (model.UnexpectedTopicReport, error) {
			return model.UnexpectedTopicReport{
				ID:               unexpectedTopicReportEntity.ID,
				TopicID:          unexpectedTopicReportEntity.TopicID,
				LearnHoursActual: unexpectedTopicReportEntity.LearnHoursActual,
				KnowledgeAreaID:  unexpectedTopicReportEntity.KnowledgeAreaID,
			}, nil
		},
	)

	// UnexpectedTopicReport: model -> entity
	RegisterMapFunc(
		func(unexpectedTopicReportModel *model.UnexpectedTopicReport) (entity.UnexpectedTopicReport, error) {
			var knowledgeAreaPtr *entity.KnowledgeArea
			if unexpectedTopicReportModel.KnowledgeAreaID != 0 {
				knowledgeAreaPtr = &entity.KnowledgeArea{
					ID: unexpectedTopicReportModel.KnowledgeAreaID,
				}
			}

			// Map Topic with language support if available
			var topicPtr *entity.Topic
			if unexpectedTopicReportModel.Topic.ID != 0 {
				if topic, err := Map[model.Topic, entity.Topic](&unexpectedTopicReportModel.Topic); err == nil {
					topicPtr = &topic
				}
			}

			return entity.UnexpectedTopicReport{
				ID:               unexpectedTopicReportModel.ID,
				TopicID:          unexpectedTopicReportModel.TopicID,
				LearnHoursActual: unexpectedTopicReportModel.LearnHoursActual,
				KnowledgeAreaID:  unexpectedTopicReportModel.KnowledgeAreaID,
				KnowledgeArea:    knowledgeAreaPtr,
				Topic:            topicPtr,
			}, nil
		},
	)

	// UnexpectedKnowledgeAreaReport: entity -> model
	RegisterMapFunc(
		func(unexpectedKnowledgeAreaReportEntity *entity.UnexpectedKnowledgeAreaReport) (model.UnexpectedKnowledgeAreaReport, error) {
			// No incluir TopicReports en el mapeo inicial para evitar problemas de FK
			// Se guardarán por separado después
			return model.UnexpectedKnowledgeAreaReport{
				ID:              unexpectedKnowledgeAreaReportEntity.ID,
				Name:            unexpectedKnowledgeAreaReportEntity.Name,
				Lang:            unexpectedKnowledgeAreaReportEntity.Lang,
				TotalLearnHours: unexpectedKnowledgeAreaReportEntity.TotalLearnHours,
				// TopicReports se manejarán por separado
			}, nil
		},
	)

	// UnexpectedKnowledgeAreaReport: model -> entity
	RegisterMapFunc(
		func(unexpectedKnowledgeAreaReportModel *model.UnexpectedKnowledgeAreaReport) (entity.UnexpectedKnowledgeAreaReport, error) {
			var topicReports []entity.UnexpectedTopicReport
			for _, tr := range unexpectedKnowledgeAreaReportModel.TopicReports {
				trEntity, err := Map[model.UnexpectedTopicReport, entity.UnexpectedTopicReport](&tr)
				if err != nil {
					return entity.UnexpectedKnowledgeAreaReport{}, err
				}
				topicReports = append(topicReports, trEntity)
			}

			return entity.UnexpectedKnowledgeAreaReport{
				ID:              unexpectedKnowledgeAreaReportModel.ID,
				Name:            unexpectedKnowledgeAreaReportModel.Name,
				Lang:            unexpectedKnowledgeAreaReportModel.Lang,
				TotalLearnHours: unexpectedKnowledgeAreaReportModel.TotalLearnHours,
				TopicReports:    topicReports,
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

			// No incluir UnexpectedKnowledgeAreaReports en el mapeo inicial
			// Se guardarán por separado después de que el reporte principal esté guardado

			return model.Report{
				ID:                   reportEntity.ID,
				ProfessionalRoleID:   reportEntity.ProfessionalRole.ID,
				DegreeProgramID:      reportEntity.DegreeProgram.ID,
				Score:                roundScore(reportEntity.Score),
				KnowledgeAreaReports: knowledgeAreaReports,
				// UnexpectedKnowledgeAreaReports se manejarán por separado
			}, nil
		},
	)

	// Report: model -> entity
	RegisterMapFunc(
		func(reportModel *model.Report) (entity.Report, error) {
			// Mapear KnowledgeAreaReports con TopicReports
			var knowledgeAreaReports []entity.KnowledgeAreaReport
			for _, kar := range reportModel.KnowledgeAreaReports {
				// Mapear TopicReports
				var topicReports []entity.TopicReport
				for _, tr := range kar.TopicReports {
					topicReport := entity.TopicReport{
						ID:                 tr.ID,
						TopicID:            tr.TopicID,
						Name:               tr.Name,
						LearnHoursExpected: tr.LearnHoursExpected,
						LearnHoursActual:   tr.LearnHoursActual,
					}
					topicReports = append(topicReports, topicReport)
				}

				karEntity := entity.KnowledgeAreaReport{
					ID:                      kar.ID,
					Name:                    kar.Name,
					TotalLearnHoursExpected: kar.TotalLearnHoursExpected,
					TotalLearnHoursActual:   kar.TotalLearnHoursActual,
					ScoreExpected:           roundScore(kar.ScoreExpected),
					ScoreGot:                roundScore(kar.ScoreGot),
					TopicReports:            topicReports,
				}
				knowledgeAreaReports = append(knowledgeAreaReports, karEntity)
			}

			// Mapear UnexpectedKnowledgeAreaReports con UnexpectedTopicReports
			var unexpectedKnowledgeAreaReports []entity.UnexpectedKnowledgeAreaReport
			for _, ukar := range reportModel.UnexpectedKnowledgeAreaReports {
				// Mapear UnexpectedTopicReports usando el idioma específico de la base de datos
				var unexpectedTopicReports []entity.UnexpectedTopicReport
				for _, utr := range ukar.TopicReports {
					// Use the language stored in the UnexpectedKnowledgeAreaReport
					unexpectedTopicReport, err := MapUnexpectedTopicReportWithLanguage(&utr, ukar.Lang)
					if err != nil {
						return entity.Report{}, err
					}
					unexpectedTopicReports = append(unexpectedTopicReports, unexpectedTopicReport)
				}

				ukarEntity := entity.UnexpectedKnowledgeAreaReport{
					ID:              ukar.ID,
					Name:            ukar.Name,
					TotalLearnHours: ukar.TotalLearnHours,
					TopicReports:    unexpectedTopicReports,
				}
				unexpectedKnowledgeAreaReports = append(unexpectedKnowledgeAreaReports, ukarEntity)
			}

			// Mapear DurationUnit si está disponible
			var durationUnitPtr *entity.DurationUnit
			if reportModel.DegreeProgram.DurationUnit.ID != 0 {
				if du, err := Map[model.DurationUnit, entity.DurationUnit](&reportModel.DegreeProgram.DurationUnit); err == nil {
					durationUnitPtr = &du
				}
			}

			// Mapear UserCreator con Person y HigherEducationInstitution
			var userCreatorPtr *entity.User
			if reportModel.DegreeProgram.UserCreator.ID != 0 {
				var personPtr *entity.Person
				if reportModel.DegreeProgram.UserCreator.Person.ID != 0 {
					var higherEducationInstitutionPtr *entity.HigherEducationInstitution
					if reportModel.DegreeProgram.UserCreator.Person.HigherEducationInstitution.Snies != 0 {
						higherEducationInstitutionPtr = &entity.HigherEducationInstitution{
							Snies: reportModel.DegreeProgram.UserCreator.Person.HigherEducationInstitution.Snies,
							Name:  reportModel.DegreeProgram.UserCreator.Person.HigherEducationInstitution.Name,
						}
					}

					personPtr = &entity.Person{
						ID:                         reportModel.DegreeProgram.UserCreator.Person.ID,
						Name:                       reportModel.DegreeProgram.UserCreator.Person.Name,
						Lastname:                   reportModel.DegreeProgram.UserCreator.Person.Lastname,
						HigherEducationInstitution: higherEducationInstitutionPtr,
					}
				}

				userCreatorPtr = &entity.User{
					ID:     reportModel.DegreeProgram.UserCreator.ID,
					Email:  reportModel.DegreeProgram.UserCreator.Email,
					Person: personPtr,
				}
			}

			// Mapear DegreeProgram con DurationUnit y UserCreator
			degreeProgram := entity.DegreeProgram{
				ID:                  reportModel.DegreeProgram.ID,
				Snies:               reportModel.DegreeProgram.Snies,
				Name:                reportModel.DegreeProgram.Name,
				TotalCredits:        reportModel.DegreeProgram.TotalCredits,
				DurationValue:       reportModel.DegreeProgram.DurationValue,
				DurationUnitID:      reportModel.DegreeProgram.DurationUnitID,
				DurationUnit:        durationUnitPtr,
				ProgramFocus:        reportModel.DegreeProgram.ProgramFocus,
				EntryProfile:        reportModel.DegreeProgram.EntryProfile,
				GraduateProfile:     reportModel.DegreeProgram.GraduateProfile,
				ProfessionalProfile: reportModel.DegreeProgram.ProfessionalProfile,
				CreatedBy:           reportModel.DegreeProgram.CreatedBy,
				UserCreator:         userCreatorPtr,
			}

			// Mapear ProfessionalRole sin KnowledgeAreaExpected para evitar errores
			professionalRole := entity.ProfessionalRole{
				ID:   reportModel.ProfessionalRole.ID,
				Name: reportModel.ProfessionalRole.Name,
				// KnowledgeAreaExpected se deja vacío porque no se necesita para el reporte
			}

			// Extraer HigherEducationInstitution desde UserCreator.Person si está disponible
			var higherEducationInstitutionPtr *entity.HigherEducationInstitution
			if degreeProgram.UserCreator != nil &&
				degreeProgram.UserCreator.Person != nil &&
				degreeProgram.UserCreator.Person.HigherEducationInstitution != nil {
				higherEducationInstitutionPtr = degreeProgram.UserCreator.Person.HigherEducationInstitution
			}

			return entity.Report{
				ID:                             reportModel.ID,
				DegreeProgram:                  degreeProgram,
				ProfessionalRole:               professionalRole,
				KnowledgeAreaReports:           knowledgeAreaReports,
				UnexpectedKnowledgeAreaReports: unexpectedKnowledgeAreaReports,
				Score:                          roundScore(reportModel.Score),
				CreatedAt:                      reportModel.CreatedAt,
				HigherEducationInstitution:     higherEducationInstitutionPtr,
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

				// Mapear UserCreator con Person y HigherEducationInstitution
				if reportFeedbackModel.Report.DegreeProgram.UserCreator.ID != 0 {
					report.DegreeProgram.UserCreator = &entity.User{
						ID:    reportFeedbackModel.Report.DegreeProgram.UserCreator.ID,
						Email: reportFeedbackModel.Report.DegreeProgram.UserCreator.Email,
					}

					// Mapear Person si está disponible
					if reportFeedbackModel.Report.DegreeProgram.UserCreator.Person.ID != 0 {
						report.DegreeProgram.UserCreator.Person = &entity.Person{
							ID:       reportFeedbackModel.Report.DegreeProgram.UserCreator.Person.ID,
							Name:     reportFeedbackModel.Report.DegreeProgram.UserCreator.Person.Name,
							Lastname: reportFeedbackModel.Report.DegreeProgram.UserCreator.Person.Lastname,
						}

						// Mapear HigherEducationInstitution si está disponible
						if reportFeedbackModel.Report.DegreeProgram.UserCreator.Person.HigherEducationInstitution.Snies != 0 {
							report.DegreeProgram.UserCreator.Person.HigherEducationInstitution = &entity.HigherEducationInstitution{
								Snies: reportFeedbackModel.Report.DegreeProgram.UserCreator.Person.HigherEducationInstitution.Snies,
								Name:  reportFeedbackModel.Report.DegreeProgram.UserCreator.Person.HigherEducationInstitution.Name,
							}
						}
					}
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
				})
			}

			return command.CreateReportFeedbackCommand{
				ReportID:               req.ReportID,
				GeneralComments:        req.GeneralComments,
				Recommendations:        req.Recommendations,
				KnowledgeAreaFeedbacks: knowledgeAreaFeedbacks,
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
			)

			// Agregar knowledge area feedbacks
			for _, kaFeedbackCmd := range cmd.KnowledgeAreaFeedbacks {
				kaFeedback := entity.NewKnowledgeAreaFeedback(
					0, // Se asignará después de crear el feedback principal
					kaFeedbackCmd.KnowledgeAreaReportID,
					kaFeedbackCmd.Comments,
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
					ScoreExpected:           roundScore(kar.ScoreExpected),
					ScoreGot:                roundScore(kar.ScoreGot),
					TopicReports:            topicReports,
				})
			}

			// Map UnexpectedKnowledgeAreaReports
			var unexpectedKnowledgeAreaReports []response.UnexpectedKnowledgeAreaReportResponse
			for _, ukar := range reportEntity.UnexpectedKnowledgeAreaReports {
				var unexpectedTopicReports []response.UnexpectedTopicReportResponse
				for _, utr := range ukar.TopicReports {
					var topicName string
					if utr.Topic != nil {
						topicName = utr.Topic.Name
					}

					unexpectedTopicReports = append(unexpectedTopicReports, response.UnexpectedTopicReportResponse{
						ID:               utr.ID,
						TopicID:          utr.TopicID,
						Name:             topicName,
						LearnHoursActual: utr.LearnHoursActual,
						KnowledgeAreaID:  utr.KnowledgeAreaID,
					})
				}

				unexpectedKnowledgeAreaReports = append(unexpectedKnowledgeAreaReports, response.UnexpectedKnowledgeAreaReportResponse{
					ID:              ukar.ID,
					Name:            ukar.Name,
					Lang:            ukar.Lang,
					TotalLearnHours: ukar.TotalLearnHours,
					TopicReports:    unexpectedTopicReports,
				})
			}

			return response.ListReportResponse{
				ID:                             reportEntity.ID,
				DegreeProgramID:                reportEntity.DegreeProgram.ID,
				DegreeProgram:                  degreeProgramResp,
				ProfessionalRole:               professionalRoleResp,
				Score:                          roundScore(reportEntity.Score),
				CreatedAt:                      reportEntity.CreatedAt,
				KnowledgeAreaReports:           knowledgeAreaReports,
				UnexpectedKnowledgeAreaReports: unexpectedKnowledgeAreaReports,
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
				})
			}

			// Mapear Report Summary
			reportSummary := response.ReportSummaryResponse{
				ID:    reportFeedbackEntity.Report.ID,
				Score: roundScore(reportFeedbackEntity.Report.Score),
				DegreeProgram: response.DegreeProgramSummaryResponse{
					ID:    reportFeedbackEntity.Report.DegreeProgram.ID,
					Name:  reportFeedbackEntity.Report.DegreeProgram.Name,
					Snies: reportFeedbackEntity.Report.DegreeProgram.Snies,
				},
				CreatedAt: reportFeedbackEntity.Report.CreatedAt.Format("2006-01-02T15:04:05Z"),
			}

			// Mapear HigherEducationInstitution si está disponible
			if reportFeedbackEntity.Report.DegreeProgram.UserCreator != nil &&
				reportFeedbackEntity.Report.DegreeProgram.UserCreator.Person != nil &&
				reportFeedbackEntity.Report.DegreeProgram.UserCreator.Person.HigherEducationInstitution != nil {
				reportSummary.HigherEducationInstitution = response.InstitutionSummaryResponse{
					ID:   reportFeedbackEntity.Report.DegreeProgram.UserCreator.Person.HigherEducationInstitution.Snies,
					Name: reportFeedbackEntity.Report.DegreeProgram.UserCreator.Person.HigherEducationInstitution.Name,
				}
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
				Score:                      roundScore(reportEntity.Score),
				DegreeProgram:              degreeProgramResp,
				ProfessionalRole:           professionalRoleResp,
				HigherEducationInstitution: institutionResp,
				CreatedAt:                  reportEntity.CreatedAt.Format("2006-01-02T15:04:05Z"),
				HasFeedback:                false, // This will be set by the service based on business logic
			}, nil
		},
	)
}

// MapUnexpectedTopicReportWithLanguage maps UnexpectedTopicReport using the specific language stored in the database
func MapUnexpectedTopicReportWithLanguage(unexpectedTopicReportModel *model.UnexpectedTopicReport, lang string) (entity.UnexpectedTopicReport, error) {
	var knowledgeAreaPtr *entity.KnowledgeArea
	if unexpectedTopicReportModel.KnowledgeAreaID != 0 {
		knowledgeAreaPtr = &entity.KnowledgeArea{
			ID: unexpectedTopicReportModel.KnowledgeAreaID,
		}
	}

	// Map Topic with the specific language
	var topicPtr *entity.Topic
	if unexpectedTopicReportModel.Topic.ID != 0 {
		if langMapper := GetTopicMapperWithLanguage(lang); langMapper != nil {
			if topic, err := langMapper(&unexpectedTopicReportModel.Topic); err == nil {
				topicPtr = &topic
			}
		} else {
			// Fallback to default mapper
			if topic, err := Map[model.Topic, entity.Topic](&unexpectedTopicReportModel.Topic); err == nil {
				topicPtr = &topic
			}
		}
	}

	return entity.UnexpectedTopicReport{
		ID:               unexpectedTopicReportModel.ID,
		TopicID:          unexpectedTopicReportModel.TopicID,
		LearnHoursActual: unexpectedTopicReportModel.LearnHoursActual,
		KnowledgeAreaID:  unexpectedTopicReportModel.KnowledgeAreaID,
		KnowledgeArea:    knowledgeAreaPtr,
		Topic:            topicPtr,
	}, nil
}
