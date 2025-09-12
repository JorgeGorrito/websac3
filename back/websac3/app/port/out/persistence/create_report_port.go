package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type CreateReportPort interface {
	Create(report *entity.Report, ctx db.Context) (*entity.Report, error)
}
