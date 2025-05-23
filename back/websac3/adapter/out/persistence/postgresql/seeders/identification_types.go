package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db/postgres"
	"websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/port/out/persistence"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_IDENTIFICATION_TYPE_SEED = "adapter/out/persistence/seeders/seeds/identification_types.json"
)

type identificationTypes struct{}

func IdentificationTypes() Seeder {
	return &identificationTypes{}
}

func (r *identificationTypes) Seed(tx persistence.Transaction) error {
	var pgTx *postgres.Transaction = tx.(*postgres.Transaction)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []models.IdentificationType = make([]models.IdentificationType, 0)
	if err := decoder.Decode(DEFAULT_PATH_IDENTIFICATION_TYPE_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, identificationType := range dataToSeed {
		if err := pgTx.Tx().Create(&identificationType).Error; err != nil {
			return err
		}
	}
	return nil
}
