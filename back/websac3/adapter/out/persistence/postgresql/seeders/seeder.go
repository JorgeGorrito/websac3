package seeders

import (
	"errors"
	"fmt"
	"websac3/adapter/out/persistence/postgresql/db"
	_db "websac3/app/port/out/persistence/db"
)

type Seeder interface {
	Seed(_db.Context) error
}

type NewSeeder func() Seeder

var registry map[string]NewSeeder = map[string]NewSeeder{
	"essential_data": EssentialData,
}

func GetSeederConstructorByName(name string) NewSeeder {
	return registry[name]
}

func ResetAutoIncrement(ctx _db.Context, tableName string) error {
	dbCtx, ok := ctx.(*db.Context)
	if !ok {
		return errors.New("invalid db context type")
	}
	sql := fmt.Sprintf(`
		SELECT setval(
			pg_get_serial_sequence('%s', 'id'),
			COALESCE(MAX(id), 0) + 1,
			false
		) FROM %s;
	`, tableName, tableName)

	return dbCtx.DB().Exec(sql).Error
}
