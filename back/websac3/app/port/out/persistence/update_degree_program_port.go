package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type UpdateDegreeProgramPort interface {
	Update(degreeProgram *entity.DegreeProgram, ctx db.Context) error
}
