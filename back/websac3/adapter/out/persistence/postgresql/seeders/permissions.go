package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_PERMISSION_SEED = "adapter/out/persistence/postgresql/seeders/seeds/permissions.json"
)

type permissions struct{}

func Permissions() Seeder {
	return &permissions{}
}

func (p *permissions) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []model.Permission = make([]model.Permission, 0)

	if err := decoder.Decode(DEFAULT_PATH_PERMISSION_SEED, &dataToSeed); err != nil {
		return err
	}

	for _, permission := range dataToSeed {
		if err := dbCtx.DB().Create(&permission).Error; err != nil {
			return err
		}
	}
	return nil
}
