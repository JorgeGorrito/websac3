package request

type AcceptExpertConsultationRequest struct {
	ExpertResponse string `json:"expert_response" binding:"required"`
}
