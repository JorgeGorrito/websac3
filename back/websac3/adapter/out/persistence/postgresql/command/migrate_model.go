package command

import (
	"fmt"
	"maps"
	"slices"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/port/out/persistence"
	"websac3/common/dependencies/container"
	"websac3/common/validator"

	"github.com/JorgeGorrito/anise-with-gin/anise/command"
)

type MigrateModel struct {
	cmdprinter         command.CMDPrinter
	modelNameToMigrate string
	paramsReceived     []string
	dbManager          *db.Manager
}

func NewMigrateModel(params map[string]string, cmdprinter command.CMDPrinter) command.Command {
	var modelNameReceived string = params["model"]
	var paramsReceived []string = slices.Collect(maps.Keys(params))
	return &MigrateModel{
		cmdprinter:         cmdprinter,
		modelNameToMigrate: modelNameReceived,
		paramsReceived:     paramsReceived,
		dbManager: func() *db.Manager {
			var dbManager *db.Manager = container.Inject[persistence.Manager]().(*db.Manager)
			return dbManager
		}(),
	}
}

func (m *MigrateModel) Execute() error {
	var err error = m.dbManager.ExecuteInTransaction(func(ctx persistence.Context) error {
		var dbCtx *db.Context = ctx.(*db.Context)
		if err := validator.ValidateParamsRequired(m.paramsReceived, []string{"model"}); err != nil {
			return err
		}
		constructor := model.GetConstructModelBaseByName(m.modelNameToMigrate)
		if constructor == nil {
			return fmt.Errorf("model %s not found", m.modelNameToMigrate)
		}

		modelToMigrate := constructor()
		if err := dbCtx.
			DB().
			AutoMigrate(modelToMigrate); err != nil {
			return err
		}
		return nil
	})
	return err
}
