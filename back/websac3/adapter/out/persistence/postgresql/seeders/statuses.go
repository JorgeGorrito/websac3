package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_ACCESS_REQUEST_STATUS_SEED = "adapter/out/persistence/postgresql/seeders/seeds/statuses.json"
)

type accessRequestStatuses struct{}

func AccessRequestStatuses() Seeder {
	return &accessRequestStatuses{}
}

func (a *accessRequestStatuses) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var dataToSeed []model.Status = make([]model.Status, 0)
	if err := decoder.Decode(DEFAULT_PATH_ACCESS_REQUEST_STATUS_SEED, &dataToSeed); err != nil {
		return err
	}
	for _, accessRequestStatus := range dataToSeed {
		if err := dbCtx.DB().Create(&accessRequestStatus).Error; err != nil {
			return err
		}
	}
	return nil
}
