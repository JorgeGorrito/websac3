package response

// ListExpertConsultationStatusResponse representa la respuesta para un estado de asesoría de experto
type ListExpertConsultationStatusResponse struct {
	ID   uint   `json:"id" example:"1"`
	Name string `json:"name" example:"Pendiente"`
}
