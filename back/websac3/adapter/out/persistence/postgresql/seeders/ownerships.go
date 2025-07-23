package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/port/out/persistence"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_OWNERSHIP_SEED = "adapter/out/persistence/postgresql/seeders/seeds/ownerships.json"
)

type ownerships struct{}

func Ownerships() Seeder {
	return &ownerships{}
}

func (o *ownerships) Seed(ctx persistence.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []model.Ownership = make([]model.Ownership, 0)

	if err := decoder.Decode(DEFAULT_PATH_OWNERSHIP_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, ownership := range dataToSeed {
		if err := dbCtx.DB().Create(&ownership).Error; err != nil {
			return err
		}
	}
	return nil
}
