package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_COURSE_NATURES_SEED         = "adapter/out/persistence/postgresql/seeders/seeds/course_natures.json"
	DEFAULT_PATH_COURSE_NATURE_NAMES_EN_SEED = "adapter/out/persistence/postgresql/seeders/seeds/course_nature_names_en.json"
	DEFAULT_PATH_COURSE_NATURE_NAMES_ES_SEED = "adapter/out/persistence/postgresql/seeders/seeds/course_nature_names_es.json"
)

type courseNatures struct{}

func CourseNatures() Seeder {
	return &courseNatures{}
}

func (s *courseNatures) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var courseNaturesToSeed []model.CourseNature = make([]model.CourseNature, 0)

	if err := decoder.Decode(DEFAULT_PATH_COURSE_NATURES_SEED, &courseNaturesToSeed); err != nil {
		return err
	}

	for _, courseNature := range courseNaturesToSeed {
		if err := dbCtx.DB().Create(&courseNature).Error; err != nil {
			return err
		}
	}

	var courseNatureNamesToSeed []model.CourseNatureName = make([]model.CourseNatureName, 0)
	var courseNatureNamesEsToSeed []model.CourseNatureName = make([]model.CourseNatureName, 0)
	var courseNatureNamesEnToSeed []model.CourseNatureName = make([]model.CourseNatureName, 0)

	if err := decoder.Decode(DEFAULT_PATH_COURSE_NATURE_NAMES_ES_SEED, &courseNatureNamesEsToSeed); err != nil {
		return err
	}

	if err := decoder.Decode(DEFAULT_PATH_COURSE_NATURE_NAMES_EN_SEED, &courseNatureNamesEnToSeed); err != nil {
		return err
	}

	courseNatureNamesToSeed = append(courseNatureNamesToSeed, courseNatureNamesEsToSeed...)
	courseNatureNamesToSeed = append(courseNatureNamesToSeed, courseNatureNamesEnToSeed...)

	for _, courseNatureName := range courseNatureNamesToSeed {
		if err := dbCtx.DB().Create(&courseNatureName).Error; err != nil {
			return err
		}
	}

	ResetAutoIncrement(dbCtx, "course_natures")
	ResetAutoIncrement(dbCtx, "course_nature_names")

	return nil
}
