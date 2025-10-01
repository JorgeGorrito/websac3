package mapper

import (
	"websac3/adapter/in/web/response"
	"websac3/app/domain/entity"
)

func registerStatisticsMappers() {
	RegisterMapFunc(func(stats *entity.UserStatistics) (response.GetUserStatisticsResponse, error) {
		statsResponse := response.GetUserStatisticsResponse{
			Role: stats.Role,
		}

		// Mapear estadísticas de Program Lead si existen
		if stats.ProgramLeadStats != nil {
			statsResponse.ProgramLeadStats = &response.ProgramLeadStatisticsResponse{
				RegisteredPrograms:  stats.ProgramLeadStats.RegisteredPrograms,
				GeneratedReports:    stats.ProgramLeadStats.GeneratedReports,
				ActiveConsultations: stats.ProgramLeadStats.ActiveConsultations,
			}
		}

		// Mapear estadísticas de Cybersecurity Auditor si existen
		if stats.CybersecurityAuditorStats != nil {
			statsResponse.CybersecurityAuditorStats = &response.CybersecurityAuditorStatisticsResponse{
				FeedbackReports:      stats.CybersecurityAuditorStats.FeedbackReports,
				PendingReports:       stats.CybersecurityAuditorStats.PendingReports,
				ConsultationRequests: stats.CybersecurityAuditorStats.ConsultationRequests,
			}
		}

		// Mapear estadísticas de Admin si existen
		if stats.AdminStats != nil {
			statsResponse.AdminStats = &response.AdminStatisticsResponse{
				ActiveUsers:         stats.AdminStats.ActiveUsers,
				TotalReports:        stats.AdminStats.TotalReports,
				TotalConsultations:  stats.AdminStats.TotalConsultations,
				TotalAccessRequests: stats.AdminStats.TotalAccessRequests,
			}
		}

		return statsResponse, nil
	})
}
