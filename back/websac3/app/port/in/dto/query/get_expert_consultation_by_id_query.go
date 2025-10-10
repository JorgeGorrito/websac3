package query

type GetExpertConsultationByIDQuery struct {
	ConsultationID uint     `json:"-" validate:"required,gt=0"`
	UserID         uint     `json:"-"`
	Permissions    []string `json:"permissions"`
}

