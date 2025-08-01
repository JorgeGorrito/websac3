package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/port/out/persistence"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_ACTION_SEED = "adapter/out/persistence/postgresql/seeders/seeds/actions.json"
)

type actions struct{}

func Actions() Seeder {
	return &actions{}
}

func (a *actions) Seed(ctx persistence.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []model.Action = make([]model.Action, 0)

	if err := decoder.Decode(DEFAULT_PATH_ACTION_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, permission := range dataToSeed {
		if err := dbCtx.DB().Create(&permission).Error; err != nil {
			return err
		}
	}
	return nil
}
