package query

import (
	"websac3/common/filter"
	"websac3/common/paginator"
)

// ListExpertConsultationStatusQuery representa la consulta para listar estados de asesorías de experto
type ListExpertConsultationStatusQuery struct {
	PaginationParams paginator.PaginationParams `json:"pagination_params"`
	Filters          filter.Params              `json:"filters"`
}
