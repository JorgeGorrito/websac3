package seeders

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/decoder"
)

const (
	DEFAULT_PATH_PEOPLE_SEED = "adapter/out/persistence/postgresql/seeders/seeds/people.json"
	DEFAULT_PATH_USER_SEED   = "adapter/out/persistence/postgresql/seeders/seeds/users.json"
)

type defaultUsers struct{}

func DefaultUsers() Seeder {
	return &defaultUsers{}
}

func (a *defaultUsers) Seed(ctx _db.Context) error {
	var dbCtx *db.Context = ctx.(*db.Context)
	var decoder decoder.Decoder = decoder.Json()
	var usersToSeed []model.User = make([]model.User, 0)
	var peopleToSeed []model.Person = make([]model.Person, 0)

	if err := decoder.Decode(DEFAULT_PATH_PEOPLE_SEED, &peopleToSeed); err != nil {
		return err
	}
	if err := decoder.Decode(DEFAULT_PATH_USER_SEED, &usersToSeed); err != nil {
		return err
	}

	for _, person := range peopleToSeed {
		if err := dbCtx.DB().Create(&person).Error; err != nil {
			return err
		}
	}
	for _, user := range usersToSeed {
		if err := dbCtx.DB().Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}
