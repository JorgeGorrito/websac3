package command

import (
	"fmt"
	"maps"
	"slices"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/dependencies/container"
	"websac3/common/validator"

	"github.com/JorgeGorrito/anise-with-gin/anise/command"
)

type MigrateModel struct {
	cmdprinter         command.CMDPrinter
	validator          validator.Validator
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
		validator:          container.Inject[validator.Validator](),
		dbManager: func() *db.Manager {
			var dbManager *db.Manager = container.Inject[_db.Manager]().(*db.Manager)
			return dbManager
		}(),
	}
}

func (m *MigrateModel) Execute() error {
	var err error = m.dbManager.ExecuteInTransaction(func(ctx _db.Context) error {
		var dbCtx *db.Context = ctx.(*db.Context)
		if err := m.validator.ValidateParamsRequired(m.paramsReceived, []string{"model"}, "es"); err != nil {
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
