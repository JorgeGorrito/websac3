package seeders

import "websac3/app/port/out/persistence/db"

type Seeder interface {
	Seed(db.Context) error
}

type NewSeeder func() Seeder

var registry map[string]NewSeeder = map[string]NewSeeder{
	"essential_data": EssentialData,
}

func GetSeederConstructorByName(name string) NewSeeder {
	return registry[name]
}
