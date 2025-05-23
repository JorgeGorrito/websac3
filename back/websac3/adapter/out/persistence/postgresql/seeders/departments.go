package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db/postgres"
	"websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/port/out/persistence"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_DEPARTMENT_SEED = "adapter/out/persistence/seeders/seeds/departments.json"
)

type departments struct{}

func Departments() Seeder {
	return &departments{}
}

func (d *departments) Seed(tx persistence.Transaction) error {
	var pgTx *postgres.Transaction = tx.(*postgres.Transaction)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []models.Department = make([]models.Department, 0)
	if err := decoder.Decode(DEFAULT_PATH_DEPARTMENT_SEED, &dataToSeed); err != nil {
		return nil
	}
	for _, department := range dataToSeed {
		if err := pgTx.Tx().Create(&department).Error; err != nil {
			return err
		}
	}
	return nil
}
