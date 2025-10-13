package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_COURSE_TYPES_SEED         = "adapter/out/persistence/postgresql/seeders/seeds/course_types.json"
	DEFAULT_PATH_COURSE_TYPE_NAMES_EN_SEED = "adapter/out/persistence/postgresql/seeders/seeds/course_type_names_en.json"
	DEFAULT_PATH_COURSE_TYPE_NAMES_ES_SEED = "adapter/out/persistence/postgresql/seeders/seeds/course_type_names_es.json"
)

type courseTypes struct{}

func CourseTypes() Seeder {
	return &courseTypes{}
}

func (s *courseTypes) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var courseTypesToSeed []model.CourseType = make([]model.CourseType, 0)

	if err := decoder.Decode(DEFAULT_PATH_COURSE_TYPES_SEED, &courseTypesToSeed); err != nil {
		return err
	}

	for _, courseType := range courseTypesToSeed {
		if err := dbCtx.DB().Create(&courseType).Error; err != nil {
			return err
		}
	}

	var courseTypeNamesToSeed []model.CourseTypeName = make([]model.CourseTypeName, 0)
	var courseTypeNamesEsToSeed []model.CourseTypeName = make([]model.CourseTypeName, 0)
	var courseTypeNamesEnToSeed []model.CourseTypeName = make([]model.CourseTypeName, 0)

	if err := decoder.Decode(DEFAULT_PATH_COURSE_TYPE_NAMES_ES_SEED, &courseTypeNamesEsToSeed); err != nil {
		return err
	}

	if err := decoder.Decode(DEFAULT_PATH_COURSE_TYPE_NAMES_EN_SEED, &courseTypeNamesEnToSeed); err != nil {
		return err
	}

	courseTypeNamesToSeed = append(courseTypeNamesToSeed, courseTypeNamesEsToSeed...)
	courseTypeNamesToSeed = append(courseTypeNamesToSeed, courseTypeNamesEnToSeed...)

	for _, courseTypeName := range courseTypeNamesToSeed {
		if err := dbCtx.DB().Create(&courseTypeName).Error; err != nil {
			return err
		}
	}

	ResetAutoIncrement(dbCtx, "course_types")
	ResetAutoIncrement(dbCtx, "course_type_names")

	return nil
}
