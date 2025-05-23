package command

import (
	"fmt"
	"maps"
	"slices"
	"websac3/adapter/out/persistence/postgresql/db/postgres"
	"websac3/adapter/out/persistence/postgresql/seeders"
	"websac3/app/port/out/persistence"
	"websac3/common/dependencies/container"
	"websac3/common/validator"

	"github.com/JorgeGorrito/anise-with-gin/anise/command"
)

type SeedRun struct {
	cmdprinter     command.CMDPrinter
	seedName       string
	paramsReceived []string
	txManager      *postgres.TransactionManager
}

func NewSeedRun(params map[string]string, cmdprinter command.CMDPrinter) command.Command {
	var paramsReceived []string = slices.Collect(maps.Keys(params))
	return &SeedRun{
		cmdprinter:     cmdprinter,
		seedName:       params["seed"],
		paramsReceived: paramsReceived,
		txManager: func() *postgres.TransactionManager {
			var pgTxManager *postgres.TransactionManager = container.Inject[persistence.TransactionManager]().(*postgres.TransactionManager)
			return pgTxManager
		}(),
	}
}
func (m *SeedRun) Execute() error {
	if err := validator.ValidateParamsRequired(m.paramsReceived, m.paramsReceived); err != nil {
		return err
	}

	return m.txManager.ExecuteInTransaction(func(tx persistence.Transaction) error {
		var constructor seeders.NewSeeder = seeders.GetSeederConstructorByName(m.seedName)
		if constructor == nil {
			return fmt.Errorf("seeder %s not found", m.seedName)
		}
		var seeder seeders.Seeder = constructor()
		if err := seeder.Seed(tx); err != nil {
			return err
		}
		return nil
	})
}
