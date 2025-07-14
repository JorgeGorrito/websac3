package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/port/out/persistence"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_DEPARTMENT_SEED = "adapter/out/persistence/seeders/seeds/departments.json"
)

type departments struct{}

func Departments() Seeder {
	return &departments{}
}

func (d *departments) Seed(ctx persistence.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []model.Department = make([]model.Department, 0)
	if err := decoder.Decode(DEFAULT_PATH_DEPARTMENT_SEED, &dataToSeed); err != nil {
		return nil
	}
	for _, department := range dataToSeed {
		if err := dbCtx.DB().Create(&department).Error; err != nil {
			return err
		}
	}
	return nil
}
