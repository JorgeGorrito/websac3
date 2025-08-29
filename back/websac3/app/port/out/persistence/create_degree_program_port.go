package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type CreateDegreeProgramPort interface {
	Create(degreeProgram *entity.DegreeProgram, db db.Context) error
}
