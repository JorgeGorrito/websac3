package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetDurationUnitPort interface {
	GetByNameAndLang(page, perPage uint, name string, lang string, db db.Context) ([]entity.DurationUnit, int64, error)
}
