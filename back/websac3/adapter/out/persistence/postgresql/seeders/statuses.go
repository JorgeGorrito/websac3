package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db/postgres"
	"websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/port/out/persistence"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_ACCESS_REQUEST_STATUS_SEED = "adapter/out/persistence/seeders/seeds/statuses.json"
)

type accessRequestStatuses struct{}

func AccessRequestStatuses() Seeder {
	return &accessRequestStatuses{}
}

func (a *accessRequestStatuses) Seed(tx persistence.Transaction) error {
	var pgTx *postgres.Transaction = tx.(*postgres.Transaction)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []models.Status = make([]models.Status, 0)
	if err := decoder.Decode(DEFAULT_PATH_ACCESS_REQUEST_STATUS_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, accessRequestStatus := range dataToSeed {
		if err := pgTx.Tx().Create(&accessRequestStatus).Error; err != nil {
			return err
		}
	}
	return nil
}
