package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetFormationLevelPort interface {
	GetAll(
		page uint,
		perPage uint,
		name string,
		ctx db.Context,
	) ([]entity.FormationLevel, int64, error)

	GetAllWithLang(
		page uint,
		perPage uint,
		name string,
		lang string,
		ctx db.Context,
	) ([]entity.FormationLevel, int64, error)
}
