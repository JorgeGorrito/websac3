package repository

import (
	psqlfilter "websac3/adapter/out/persistence/postgresql/filter"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/app/port/out/persistence/filter"
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

func (r *DegreeProgramRepository) GetByID(id uint, ctx _db.Context) (entity.DegreeProgram, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return entity.DegreeProgram{}, err
	}

	var dp model.DegreeProgram
	if err := dbCtx.DB().
		Model(&model.DegreeProgram{}).
		Preload("DurationUnit").
		Preload("DurationUnit.Names").
		Preload("UserCreator").
		Preload("UserCreator.Person").
		Preload("UserCreator.Person.HigherEducationInstitution").
		Preload("UserCreator.Person.HigherEducationInstitution.Ownership").
		Preload("UserCreator.Person.HigherEducationInstitution.InstitutionalCategory").
		Preload("UserCreator.Person.HigherEducationInstitution.Municipality").
		Preload("UserCreator.Person.HigherEducationInstitution.Department").
		Preload("Courses").
		Preload("Courses.Nature").
		Preload("Courses.Type").
		Preload("Courses.CourseTopics").
		Preload("Courses.CourseTopics.Topic").
		Preload("Courses.CourseTopics.Topic.Names").
		Where("degree_programs.id = ?", id).
		First(&dp).Error; err != nil {
		return entity.DegreeProgram{}, err
	}

	return mapper.Map[model.DegreeProgram, entity.DegreeProgram](&dp)
}

func (r *DegreeProgramRepository) GetByFilters(page, perPage uint, filters filter.Filters, ctx _db.Context) ([]entity.DegreeProgram, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	baseQuery := dbCtx.DB().
		Model(&model.DegreeProgram{}).
		Preload("DurationUnit").
		Preload("DurationUnit.Names").
		Preload("UserCreator").
		Preload("UserCreator.Person").
		Preload("UserCreator.Person.HigherEducationInstitution")

	dbCtx.DBSet(baseQuery)

	dbCtxFiltered, err := filters.Apply(dbCtx, psqlfilter.FiltersRegistry)
	if err != nil {
		return nil, 0, err
	}

	dbCtx, err = r.CastDbContext(dbCtxFiltered)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	if err := dbCtx.DB().Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query := dbCtx.DB().
		Offset(int((page - 1) * perPage)).
		Limit(int(perPage))

	dbCtx.DBSet(query)

	var degreeProgramsFound []model.DegreeProgram
	if err := dbCtx.DB().Find(&degreeProgramsFound).Error; err != nil {
		return nil, 0, err
	}

	var degreePrograms []entity.DegreeProgram
	for _, degreeProgram := range degreeProgramsFound {
		var degreeProgramEntity entity.DegreeProgram
		if degreeProgramEntity, err = mapper.Map[model.DegreeProgram, entity.DegreeProgram](&degreeProgram); err != nil {
			return nil, 0, err
		}
		degreePrograms = append(degreePrograms, degreeProgramEntity)
	}

	return degreePrograms, total, nil
}

func (r *DegreeProgramRepository) GetByUserIDAndFilters(page, perPage uint, userID uint, filters filter.Filters, ctx _db.Context) ([]entity.DegreeProgram, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	baseQuery := dbCtx.DB().
		Model(&model.DegreeProgram{}).
		Preload("DurationUnit").
		Preload("DurationUnit.Names").
		Preload("UserCreator").
		Preload("UserCreator.Person").
		Preload("UserCreator.Person.HigherEducationInstitution").
		Where("created_by = ?", userID)

	dbCtx.DBSet(baseQuery)

	dbCtxFiltered, err := filters.Apply(dbCtx, psqlfilter.FiltersRegistry)
	if err != nil {
		return nil, 0, err
	}

	dbCtx, err = r.CastDbContext(dbCtxFiltered)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	if err := dbCtx.DB().Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query := dbCtx.DB().
		Offset(int((page - 1) * perPage)).
		Limit(int(perPage))

	dbCtx.DBSet(query)

	var degreeProgramsFound []model.DegreeProgram
	if err := dbCtx.DB().Find(&degreeProgramsFound).Error; err != nil {
		return nil, 0, err
	}

	var degreePrograms []entity.DegreeProgram
	for _, degreeProgram := range degreeProgramsFound {
		var degreeProgramEntity entity.DegreeProgram
		if degreeProgramEntity, err = mapper.Map[model.DegreeProgram, entity.DegreeProgram](&degreeProgram); err != nil {
			return nil, 0, err
		}
		degreePrograms = append(degreePrograms, degreeProgramEntity)
	}

	return degreePrograms, total, nil
}
