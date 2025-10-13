package persistence

import (
	"websac3/app/port/out/persistence/db"
)

type DeleteDegreeProgramPort interface {
	DeleteByID(degreeProgramID uint, ctx db.Context) error
}



