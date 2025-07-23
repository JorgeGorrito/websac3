package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/port/out/persistence"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_MODULES_SEED = "adapter/out/persistence/postgresql/seeders/seeds/modules.json"
)

type modules struct{}

func Modules() Seeder {
	return &modules{}
}

func (m *modules) Seed(ctx persistence.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []model.Module = make([]model.Module, 0)

	if err := decoder.Decode(DEFAULT_PATH_MODULES_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, module := range dataToSeed {
		if err := dbCtx.DB().Create(&module).Error; err != nil {
			return err
		}
	}
	return nil
}
