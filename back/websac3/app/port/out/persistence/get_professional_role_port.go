package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetProfessionalRolePort interface {
	GetByID(id uint, lang string, db db.Context) (entity.ProfessionalRole, error)
}
