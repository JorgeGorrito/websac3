package command

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/models"
	"websac3/app/port/out/persistence"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-with-gin/anise/command"
	"gorm.io/gorm"
)

type MigrateAllModels struct {
	cmdprinter command.CMDPrinter
	db         *gorm.DB
	models     []any
	txManager  *db.TransactionManager
}

func NewMigrateAllModels(params map[string]string, cmdprinter command.CMDPrinter) command.Command {
	return &MigrateAllModels{
		cmdprinter: cmdprinter,
		db:         nil,
		models: func() (modelsToMigrate []any) {
			var modelsConstructors map[string]models.NewBaseModel = models.GetRegistryAllConstructModelBase()
			for _, constructor := range modelsConstructors {
				modelsToMigrate = append(modelsToMigrate, constructor())
			}
			return modelsToMigrate
		}(),
		txManager: func() *db.TransactionManager {
			var txm *db.TransactionManager = container.Inject[persistence.TransactionManager]().(*db.TransactionManager)
			return txm
		}(),
	}
}

func (m *MigrateAllModels) Execute() error {
	var err error = m.txManager.ExecuteInTransaction(func(tx persistence.Transaction) error {
		var pgTx *db.Transaction = tx.(*db.Transaction)
		if err := pgTx.Tx().AutoMigrate(m.models...); err != nil {
			return err
		}
		return nil
	})
	return err
}
