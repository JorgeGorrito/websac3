package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/port/out/persistence"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_ROLE_SEED = "adapter/out/persistence/seeders/seeds/roles.json"
)

type roles struct{}

func Roles() Seeder {
	return &roles{}
}

func (r *roles) Seed(ctx persistence.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []model.Role = make([]model.Role, 0)

	if err := decoder.Decode(DEFAULT_PATH_ROLE_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, role := range dataToSeed {
		if err := dbCtx.DB().Create(&role).Error; err != nil {
			return err
		}
	}
	return nil
}
