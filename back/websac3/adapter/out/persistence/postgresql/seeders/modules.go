package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_MODULES_SEED = "adapter/out/persistence/postgresql/seeders/seeds/modules.json"
)

type modules struct{}

func Modules() Seeder {
	return &modules{}
}

func (m *modules) Seed(ctx _db.Context) error {
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

	ResetAutoIncrement(dbCtx, "modules")

	return nil
}
