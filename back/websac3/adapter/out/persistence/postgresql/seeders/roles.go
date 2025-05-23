package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/models"
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

func (r *roles) Seed(tx persistence.Transaction) error {
	var pgTx *db.Transaction = tx.(*db.Transaction)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []models.Role = make([]models.Role, 0)

	if err := decoder.Decode(DEFAULT_PATH_ROLE_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, role := range dataToSeed {
		if err := pgTx.Tx().Create(&role).Error; err != nil {
			return err
		}
	}
	return nil
}
