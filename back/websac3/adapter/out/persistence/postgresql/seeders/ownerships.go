package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db/postgres"
	"websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/port/out/persistence"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_OWNERSHIP_SEED = "adapter/out/persistence/seeders/seeds/ownerships.json"
)

type ownerships struct{}

func Ownerships() Seeder {
	return &ownerships{}
}

func (o *ownerships) Seed(tx persistence.Transaction) error {
	var pgTx *postgres.Transaction = tx.(*postgres.Transaction)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []models.Ownership = make([]models.Ownership, 0)

	if err := decoder.Decode(DEFAULT_PATH_OWNERSHIP_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, ownership := range dataToSeed {
		if err := pgTx.Tx().Create(&ownership).Error; err != nil {
			return err
		}
	}
	return nil
}
