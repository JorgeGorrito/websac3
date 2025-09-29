package command

type AcceptExpertConsultationCommand struct {
	ConsultationID uint   `json:"consultation_id" validate:"required"`
	ExpertResponse string `json:"expert_response" validate:"required"`
	ExpertID       uint   `json:"-"`
}
