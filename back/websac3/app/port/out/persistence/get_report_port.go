package persistence

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/persistence/db"
)

type GetReportPort interface {
	GetByDegreeProgramID(degreeProgramID uint, page, perPage uint, ctx db.Context) ([]entity.Report, int64, error)
	GetByID(reportID uint, ctx db.Context) (entity.Report, error)
}
