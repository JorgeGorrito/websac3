package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type DegreeProgramRepository struct {
	Repository
}

func NewDegreeProgramRepository() *DegreeProgramRepository {
	return &DegreeProgramRepository{}
}

func (r *DegreeProgramRepository) Create(degreeProgramToSave *entity.DegreeProgram, ctx _db.Context) error {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return err
	}

	var degreeProgram model.DegreeProgram
	if degreeProgram, err = mapper.Map[entity.DegreeProgram, model.DegreeProgram](degreeProgramToSave); err != nil {
		return err
	}

	if err := dbCtx.DB().
		Create(&degreeProgram).
		Error; err != nil {
		return err
	}
	degreeProgramToSave.ID = degreeProgram.ID

	return nil
}
