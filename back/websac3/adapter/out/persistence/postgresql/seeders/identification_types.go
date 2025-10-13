package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_IDENTIFICATION_TYPE_SEED = "adapter/out/persistence/postgresql/seeders/seeds/identification_types.json"
)

type identificationTypes struct{}

func IdentificationTypes() Seeder {
	return &identificationTypes{}
}

func (r *identificationTypes) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []model.IdentificationType = make([]model.IdentificationType, 0)
	if err := decoder.Decode(DEFAULT_PATH_IDENTIFICATION_TYPE_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, identificationType := range dataToSeed {
		if err := dbCtx.DB().Create(&identificationType).Error; err != nil {
			return err
		}
	}

	ResetAutoIncrement(dbCtx, "identification_types")

	return nil
}
