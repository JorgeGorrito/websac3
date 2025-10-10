package repository

import (
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/mapper"
)

type DurationUnitRepository struct {
	Repository
}

func NewDurationUnitRepository() *DurationUnitRepository {
	return &DurationUnitRepository{}
}

func (r *DurationUnitRepository) GetByNameAndLang(
	page uint,
	perPage uint,
	name string,
	lang string,
	ctx _db.Context,
) ([]entity.DurationUnit, int64, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	base := dbCtx.DB().
		Model(&model.DurationUnit{})

	sub := dbCtx.DB().Table("duration_unit_names dun").
		Where("dun.duration_unit_id = duration_units.id").
		Where("dun.lang = ? AND dun.name ILIKE ?", lang, "%"+name+"%")

	base = base.
		Where("EXISTS (?)", sub).
		Preload("Names", "lang = ? AND name ILIKE ?", lang, "%"+name+"%")

	dbCtx.DBSet(base)

	var count int64
	if err := dbCtx.DB().Count(&count).Error; err != nil {
		return nil, 0, err
	}

	query := dbCtx.DB().
		Offset(int((page - 1) * perPage)).
		Limit(int(perPage))

	dbCtx.DBSet(query)

	var durationUnits []model.DurationUnit
	if err := dbCtx.DB().Find(&durationUnits).Error; err != nil {
		return nil, 0, err
	}

	var results []entity.DurationUnit
	for _, du := range durationUnits {
		durationUnitEntity, err := mapper.Map[model.DurationUnit, entity.DurationUnit](&du)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, durationUnitEntity)
	}

	return results, count, nil
}

func (r *DurationUnitRepository) GetByID(
	id uint,
	ctx _db.Context,
) (entity.DurationUnit, error) {
	dbCtx, err := r.CastDbContext(ctx)
	if err != nil {
		return entity.DurationUnit{}, err
	}

	var durationUnit model.DurationUnit
	if err := dbCtx.DB().
		Preload("Names").
		First(&durationUnit, id).Error; err != nil {
		return entity.DurationUnit{}, err
	}

	return mapper.Map[model.DurationUnit, entity.DurationUnit](&durationUnit)
}
