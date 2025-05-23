package seeders

import (
	"websac3/app/port/out/persistence"
)

type Seeder interface {
	Seed(persistence.Transaction) error
}

type NewSeeder func() Seeder

var registry map[string]NewSeeder = map[string]NewSeeder{
	"essential_data": EssentialData,
}

func GetSeederConstructorByName(name string) NewSeeder {
	return registry[name]
}
