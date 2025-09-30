package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
	"websac3/common/filter"
	"websac3/common/paginator"
)

// GetExpertConsultationStatusPort define el puerto para obtener estados de asesorías de experto
type GetExpertConsultationStatusPort interface {
	// GetAll obtiene todos los estados de asesorías de experto con paginación y filtros
	GetAll(paginationParams paginator.PaginationParams, filters filter.Params, ctx db.Context) ([]entity.ExpertConsultationStatus, uint, error)
}
