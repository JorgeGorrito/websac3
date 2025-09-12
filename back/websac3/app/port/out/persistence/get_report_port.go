package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetReportPort interface {
	GetByDegreeProgramID(degreeProgramID uint, ctx db.Context) ([]entity.Report, error)
	GetByID(reportID uint, ctx db.Context) (entity.Report, error)
}
