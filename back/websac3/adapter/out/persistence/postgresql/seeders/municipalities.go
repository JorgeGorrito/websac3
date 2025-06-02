package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
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
	var pgTx *db.Transaction = tx.(*db.Transaction)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []model.Municipality = make([]model.Municipality, 0)
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
