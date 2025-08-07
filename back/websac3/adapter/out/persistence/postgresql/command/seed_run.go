package command

import (
	"fmt"
	"maps"
	"slices"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/seeders"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/dependencies/container"
	"websac3/common/validator"

	"github.com/JorgeGorrito/anise-with-gin/anise/command"
)

type SeedRun struct {
	cmdprinter     command.CMDPrinter
	validator      validator.Validator
	seedName       string
	paramsReceived []string
	dbManager      *db.Manager
}

func NewSeedRun(params map[string]string, cmdprinter command.CMDPrinter) command.Command {
	var paramsReceived []string = slices.Collect(maps.Keys(params))
	return &SeedRun{
		cmdprinter:     cmdprinter,
		seedName:       params["seed"],
		paramsReceived: paramsReceived,
		validator:      container.Inject[validator.Validator](),
		dbManager: func() *db.Manager {
			var dbManager *db.Manager = container.Inject[_db.Manager]().(*db.Manager)
			return dbManager
		}(),
	}
}
func (m *SeedRun) Execute() error {
	if err := m.validator.ValidateParamsRequired(m.paramsReceived, m.paramsReceived, "es"); err != nil {
		return err
	}

	return m.dbManager.ExecuteInTransaction(func(ctx _db.Context) error {
		var constructor seeders.NewSeeder = seeders.GetSeederConstructorByName(m.seedName)
		if constructor == nil {
			return fmt.Errorf("seeder %s not found", m.seedName)
		}
		var seeder seeders.Seeder = constructor()
		if err := seeder.Seed(ctx); err != nil {
			return err
		}
		return nil
	})
}
