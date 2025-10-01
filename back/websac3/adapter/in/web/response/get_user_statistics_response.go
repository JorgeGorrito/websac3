package response

type ProgramLeadStatisticsResponse struct {
	RegisteredPrograms  int64 `json:"registered_programs"`
	GeneratedReports    int64 `json:"generated_reports"`
	ActiveConsultations int64 `json:"active_consultations"`
}

type CybersecurityAuditorStatisticsResponse struct {
	FeedbackReports      int64 `json:"feedback_reports"`
	PendingReports       int64 `json:"pending_reports"`
	ConsultationRequests int64 `json:"consultation_requests"`
}

type AdminStatisticsResponse struct {
	ActiveUsers         int64 `json:"active_users"`
	TotalReports        int64 `json:"total_reports"`
	TotalConsultations  int64 `json:"total_consultations"`
	TotalAccessRequests int64 `json:"total_access_requests"`
}

type GetUserStatisticsResponse struct {
	Role                      string                                  `json:"role"`
	ProgramLeadStats          *ProgramLeadStatisticsResponse          `json:"program_lead_stats,omitempty"`
	CybersecurityAuditorStats *CybersecurityAuditorStatisticsResponse `json:"cybersecurity_auditor_stats,omitempty"`
	AdminStats                *AdminStatisticsResponse                `json:"admin_stats,omitempty"`
}
