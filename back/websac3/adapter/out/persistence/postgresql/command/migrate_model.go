package command

import (
	"fmt"
	"maps"
	"slices"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/port/out/persistence"
	"websac3/common/dependencies/container"
	"websac3/common/validator"

	"github.com/JorgeGorrito/anise-with-gin/anise/command"
)

type MigrateModel struct {
	cmdprinter         command.CMDPrinter
	modelNameToMigrate string
	paramsReceived     []string
	txManager          *db.TransactionManager
}

func NewMigrateModel(params map[string]string, cmdprinter command.CMDPrinter) command.Command {
	var modelNameReceived string = params["model"]
	var paramsReceived []string = slices.Collect(maps.Keys(params))
	return &MigrateModel{
		cmdprinter:         cmdprinter,
		modelNameToMigrate: modelNameReceived,
		paramsReceived:     paramsReceived,
		txManager: func() *db.TransactionManager {
			var txm *db.TransactionManager = container.Inject[persistence.TransactionManager]().(*db.TransactionManager)
			return txm
		}(),
	}
}

func (m *MigrateModel) Execute() error {
	var err error = m.txManager.ExecuteInTransaction(func(tx persistence.Transaction) error {
		var pgTx *db.Transaction = tx.(*db.Transaction)
		if err := validator.ValidateParamsRequired(m.paramsReceived, []string{"model"}); err != nil {
			return err
		}
		constructor := models.GetConstructModelBaseByName(m.modelNameToMigrate)
		if constructor == nil {
			return fmt.Errorf("model %s not found", m.modelNameToMigrate)
		}

		modelToMigrate := constructor()
		if err := pgTx.Tx().AutoMigrate(modelToMigrate); err != nil {
			return err
		}
		return nil
	})
	return err
}
