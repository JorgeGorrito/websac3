package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_FORMATION_LEVEL_SEED         = "adapter/out/persistence/postgresql/seeders/seeds/formation_levels.json"
	DEFAULT_PATH_FORMATION_LEVEL_NAME_EN_SEED = "adapter/out/persistence/postgresql/seeders/seeds/formation_level_names_en.json"
	DEFAULT_PATH_FORMATION_LEVEL_NAME_ES_SEED = "adapter/out/persistence/postgresql/seeders/seeds/formation_level_names_es.json"
)

type formationLevels struct{}

func FormationLevels() Seeder {
	return &formationLevels{}
}

func (s *formationLevels) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var formationLevelsToSeed []model.FormationLevel = make([]model.FormationLevel, 0)

	if err := decoder.Decode(DEFAULT_PATH_FORMATION_LEVEL_SEED, &formationLevelsToSeed); err != nil {
		return err
	}

	for _, formationLevel := range formationLevelsToSeed {
		if err := dbCtx.DB().Create(&formationLevel).Error; err != nil {
			return err
		}
	}

	var formationLevelNamesToSeed []model.FormationLevelName = make([]model.FormationLevelName, 0)
	var formationLevelNamesEsToSeed []model.FormationLevelName = make([]model.FormationLevelName, 0)
	var formationLevelNamesEnToSeed []model.FormationLevelName = make([]model.FormationLevelName, 0)

	if err := decoder.Decode(DEFAULT_PATH_FORMATION_LEVEL_NAME_ES_SEED, &formationLevelNamesEsToSeed); err != nil {
		return err
	}

	if err := decoder.Decode(DEFAULT_PATH_FORMATION_LEVEL_NAME_EN_SEED, &formationLevelNamesEnToSeed); err != nil {
		return err
	}

	formationLevelNamesToSeed = append(formationLevelNamesToSeed, formationLevelNamesEsToSeed...)
	formationLevelNamesToSeed = append(formationLevelNamesToSeed, formationLevelNamesEnToSeed...)

	for _, formationLevelName := range formationLevelNamesToSeed {
		if err := dbCtx.DB().Create(&formationLevelName).Error; err != nil {
			return err
		}
	}

	return nil
}
