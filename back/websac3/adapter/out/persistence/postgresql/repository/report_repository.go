package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type ReportRepository struct {
	Repository
}

func NewReportRepository() *ReportRepository {
	return &ReportRepository{}
}

func (r *ReportRepository) Create(reportToSave *entity.Report, ctx _db.Context) (*entity.Report, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, err
	}

	var reportModel model.Report
	if reportModel, err = mapper.Map[entity.Report, model.Report](reportToSave); err != nil {
		return nil, err
	}

	if err := dbCtx.DB().
		Create(&reportModel).
		Error; err != nil {
		return nil, err
	}

	// Actualizar el ID en la entidad original
	reportToSave.ID = reportModel.ID

	// Devolver la entidad guardada
	return reportToSave, nil
}
