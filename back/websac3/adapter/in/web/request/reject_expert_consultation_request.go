package request

type RejectExpertConsultationRequest struct {
	ExpertResponse string `json:"expert_response" binding:"required"`
}
