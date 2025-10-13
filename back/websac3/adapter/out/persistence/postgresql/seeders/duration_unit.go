package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_DURATION_UNIT_SEED         = "adapter/out/persistence/postgresql/seeders/seeds/duration_unit.json"
	DEFAULT_PATH_DURATION_UNIT_NAME_EN_SEED = "adapter/out/persistence/postgresql/seeders/seeds/duration_unit_names_en.json"
	DEFAULT_PATH_DURATION_UNIT_NAME_ES_SEED = "adapter/out/persistence/postgresql/seeders/seeds/duration_unit_names_es.json"
)

type durationUnits struct{}

func DurationUnits() Seeder {
	return &durationUnits{}
}

func (s *durationUnits) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var durationUnitsToSeed []model.DurationUnit = make([]model.DurationUnit, 0)

	if err := decoder.Decode(DEFAULT_PATH_DURATION_UNIT_SEED, &durationUnitsToSeed); err != nil {
		return err
	}

	for _, durationUnit := range durationUnitsToSeed {
		if err := dbCtx.DB().Create(&durationUnit).Error; err != nil {
			return err
		}
	}

	var durationUnitNamesToSeed []model.DurationUnitName = make([]model.DurationUnitName, 0)
	var durationUnitNamesEsToSeed []model.DurationUnitName = make([]model.DurationUnitName, 0)
	var durationUnitNamesEnToSeed []model.DurationUnitName = make([]model.DurationUnitName, 0)

	if err := decoder.Decode(DEFAULT_PATH_DURATION_UNIT_NAME_ES_SEED, &durationUnitNamesEsToSeed); err != nil {
		return err
	}

	if err := decoder.Decode(DEFAULT_PATH_DURATION_UNIT_NAME_EN_SEED, &durationUnitNamesEnToSeed); err != nil {
		return err
	}

	durationUnitNamesToSeed = append(durationUnitNamesToSeed, durationUnitNamesEsToSeed...)
	durationUnitNamesToSeed = append(durationUnitNamesToSeed, durationUnitNamesEnToSeed...)

	for _, durationUnitName := range durationUnitNamesToSeed {
		if err := dbCtx.DB().Create(&durationUnitName).Error; err != nil {
			return err
		}
	}

	ResetAutoIncrement(dbCtx, "duration_units")
	ResetAutoIncrement(dbCtx, "duration_unit_names")

	return nil
}
