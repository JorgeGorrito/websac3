package constants

const (
	// ExpertConsultationStatusPending indica que la consulta está pendiente de respuesta
	ExpertConsultationStatusPending = "pending"

	// ExpertConsultationStatusRejected indica que la consulta ha sido rechazada por un experto
	ExpertConsultationStatusRejected = "rejected"

	// ExpertConsultationStatusAccepted indica que la consulta ha sido aceptada y respondida por un experto
	ExpertConsultationStatusAccepted = "accepted"
)

// GetValidExpertConsultationStatuses retorna todos los estados válidos para una consulta de experto
func GetValidExpertConsultationStatuses() []string {
	return []string{
		ExpertConsultationStatusPending,
		ExpertConsultationStatusRejected,
		ExpertConsultationStatusAccepted,
	}
}

// IsValidExpertConsultationStatus verifica si un estado es válido
func IsValidExpertConsultationStatus(status string) bool {
	validStatuses := GetValidExpertConsultationStatuses()
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}
