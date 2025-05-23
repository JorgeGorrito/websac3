package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db/postgres"
	"websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/port/out/persistence"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_MUNICIPALITY_SEED = "adapter/out/persistence/seeders/seeds/municipalities.json"
)

type municipalities struct{}

func Municipalities() Seeder {
	return &municipalities{}
}

func (m *municipalities) Seed(tx persistence.Transaction) error {
	var pgTx *postgres.Transaction = tx.(*postgres.Transaction)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []models.Municipality = make([]models.Municipality, 0)
	if err := decoder.Decode(DEFAULT_PATH_MUNICIPALITY_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, municipality := range dataToSeed {
		if err := pgTx.Tx().Create(&municipality).Error; err != nil {
			return err
		}
	}
	return nil
}
