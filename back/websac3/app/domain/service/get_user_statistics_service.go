package service

import (
	"websac3/app/domain/constants"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type GetUserStatisticsService struct {
	getUserPort        persistence.GetUserPort
	getStatisticsPort  persistence.GetStatisticsPort
	msgProvider        message.Provider
	persistenceManager db.Manager
}

func NewGetUserStatisticsService(
	getUserPort persistence.GetUserPort,
	getStatisticsPort persistence.GetStatisticsPort,
	msgProvider message.Provider,
	persistenceManager db.Manager,
) usecase.GetUserStatisticsUseCase {
	return &GetUserStatisticsService{
		getUserPort:        getUserPort,
		getStatisticsPort:  getStatisticsPort,
		msgProvider:        msgProvider,
		persistenceManager: persistenceManager,
	}
}

func (s *GetUserStatisticsService) Execute(userID uint, lang string) (entity.UserStatistics, error) {
	var stats entity.UserStatistics

	err := s.persistenceManager.ExecuteNonTransactional(
		func(ctx db.Context) error {
			// Obtener el usuario para conocer su rol
			user, err := s.getUserPort.GetByID(userID, ctx)
			if err != nil {
				return err
			}

			if !user.IsRegistered() {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("get_user_statistics", "user_not_found"),
				)
			}

			if user.Role == nil {
				return errs.NewValidationError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("get_user_statistics", "user_role_not_found"),
				)
			}

			stats.Role = user.Role.Name

			// Obtener estadísticas según el rol
			switch user.Role.Name {
			case constants.Guess, constants.ProgramLead:
				return s.getProgramLeadStatistics(userID, &stats, ctx)
			case constants.CybersecurityAuditor:
				return s.getCybersecurityAuditorStatistics(userID, &stats, ctx)
			case constants.Admin:
				return s.getAdminStatistics(&stats, ctx)
			default:
				return errs.NewValidationError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("get_user_statistics", "invalid_role"),
				)
			}
		},
	)

	return stats, err
}

func (s *GetUserStatisticsService) getProgramLeadStatistics(
	userID uint,
	stats *entity.UserStatistics,
	ctx db.Context,
) error {
	programLeadStats := &entity.ProgramLeadStatistics{}

	// Contar programas registrados
	count, err := s.getStatisticsPort.CountDegreeProgramsByUserID(userID, ctx)
	if err != nil {
		return err
	}
	programLeadStats.RegisteredPrograms = count

	// Contar reportes generados
	count, err = s.getStatisticsPort.CountReportsByUserID(userID, ctx)
	if err != nil {
		return err
	}
	programLeadStats.GeneratedReports = count

	// Contar asesorías activas
	count, err = s.getStatisticsPort.CountActiveConsultationsByUserID(userID, ctx)
	if err != nil {
		return err
	}
	programLeadStats.ActiveConsultations = count

	stats.ProgramLeadStats = programLeadStats
	return nil
}

func (s *GetUserStatisticsService) getCybersecurityAuditorStatistics(
	userID uint,
	stats *entity.UserStatistics,
	ctx db.Context,
) error {
	auditorStats := &entity.CybersecurityAuditorStatistics{}

	// Contar reportes retroalimentados
	count, err := s.getStatisticsPort.CountFeedbackReportsByAuditorID(userID, ctx)
	if err != nil {
		return err
	}
	auditorStats.FeedbackReports = count

	// Contar reportes pendientes
	count, err = s.getStatisticsPort.CountPendingReportsByAuditorID(userID, ctx)
	if err != nil {
		return err
	}
	auditorStats.PendingReports = count

	// Contar solicitudes de asesoría
	count, err = s.getStatisticsPort.CountConsultationRequestsByExpertID(userID, ctx)
	if err != nil {
		return err
	}
	auditorStats.ConsultationRequests = count

	stats.CybersecurityAuditorStats = auditorStats
	return nil
}

func (s *GetUserStatisticsService) getAdminStatistics(
	stats *entity.UserStatistics,
	ctx db.Context,
) error {
	adminStats := &entity.AdminStatistics{}

	// Contar usuarios activos
	count, err := s.getStatisticsPort.CountActiveUsers(ctx)
	if err != nil {
		return err
	}
	adminStats.ActiveUsers = count

	// Contar total de reportes
	count, err = s.getStatisticsPort.CountTotalReports(ctx)
	if err != nil {
		return err
	}
	adminStats.TotalReports = count

	// Contar total de consultas
	count, err = s.getStatisticsPort.CountTotalConsultations(ctx)
	if err != nil {
		return err
	}
	adminStats.TotalConsultations = count

	// Contar total de solicitudes de acceso
	count, err = s.getStatisticsPort.CountTotalAccessRequests(ctx)
	if err != nil {
		return err
	}
	adminStats.TotalAccessRequests = count

	stats.AdminStats = adminStats
	return nil
}
