package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetTopicPort interface {
	GetByNameAndLang(page, perPage uint, name string, id string, lang string, db db.Context) ([]entity.Topic, int64, error)
}
