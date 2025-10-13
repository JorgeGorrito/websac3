package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_INSTITUTIONAL_CATEGORY_SEED = "adapter/out/persistence/postgresql/seeders/seeds/institutional_categories.json"
)

type institutionalCategories struct{}

func InstitutionalCategories() Seeder {
	return &institutionalCategories{}
}

func (i *institutionalCategories) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []model.InstitutionalCategory = make([]model.InstitutionalCategory, 0)

	if err := decoder.Decode(DEFAULT_PATH_INSTITUTIONAL_CATEGORY_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, institutionalCategory := range dataToSeed {
		if err := dbCtx.DB().Create(&institutionalCategory).Error; err != nil {
			return err
		}
	}

	ResetAutoIncrement(dbCtx, "institutional_categories")

	return nil
}
