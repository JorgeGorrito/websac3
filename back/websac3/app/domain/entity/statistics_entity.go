package entity

// ProgramLeadStatistics estadísticas para director de programa o invitado
type ProgramLeadStatistics struct {
	RegisteredPrograms  int64 `json:"registered_programs"`
	GeneratedReports    int64 `json:"generated_reports"`
	ActiveConsultations int64 `json:"active_consultations"`
}

// CybersecurityAuditorStatistics estadísticas para auditor de ciberseguridad
type CybersecurityAuditorStatistics struct {
	FeedbackReports      int64 `json:"feedback_reports"`
	PendingReports       int64 `json:"pending_reports"`
	ConsultationRequests int64 `json:"consultation_requests"`
}

// AdminStatistics estadísticas para administrador
type AdminStatistics struct {
	ActiveUsers         int64 `json:"active_users"`
	TotalReports        int64 `json:"total_reports"`
	TotalConsultations  int64 `json:"total_consultations"`
	TotalAccessRequests int64 `json:"total_access_requests"`
}

// UserStatistics contenedor para todas las estadísticas según el rol
type UserStatistics struct {
	Role                      string
	ProgramLeadStats          *ProgramLeadStatistics
	CybersecurityAuditorStats *CybersecurityAuditorStatistics
	AdminStats                *AdminStatistics
}
