package usecase

import (
	"websac3/app/port/in/dto/query"
	"websac3/adapter/in/web/response"
	"websac3/common/paginator"
)

// ListExpertConsultationStatusUseCase define el caso de uso para listar estados de asesorías de experto
type ListExpertConsultationStatusUseCase interface {
	Execute(request query.ListExpertConsultationStatusQuery, lang string) (response.ApiResponse[paginator.Page[response.ListExpertConsultationStatusResponse]], error)
}
