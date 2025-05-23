package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db/postgres"
	"websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/port/out/persistence"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_INSTITUTIONAL_CATEGORY_SEED = "adapter/out/persistence/seeders/seeds/institutional_categories.json"
)

type institutionalCategories struct{}

func InstitutionalCategories() Seeder {
	return &institutionalCategories{}
}

func (i *institutionalCategories) Seed(tx persistence.Transaction) error {
	var pgTx *postgres.Transaction = tx.(*postgres.Transaction)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []models.InstitutionalCategory = make([]models.InstitutionalCategory, 0)

	if err := decoder.Decode(DEFAULT_PATH_INSTITUTIONAL_CATEGORY_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, institutionalCategory := range dataToSeed {
		if err := pgTx.Tx().Create(&institutionalCategory).Error; err != nil {
			return err
		}
	}
	return nil
}
