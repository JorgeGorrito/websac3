package command

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-with-gin/anise/command"
	"gorm.io/gorm"
)

type MigrateAllModels struct {
	cmdprinter command.CMDPrinter
	db         *gorm.DB
	models     []any
	dbManager  *db.Manager
}

func NewMigrateAllModels(params map[string]string, cmdprinter command.CMDPrinter) command.Command {
	return &MigrateAllModels{
		cmdprinter: cmdprinter,
		db:         nil,
		models: func() (modelsToMigrate []any) {
			var modelsConstructors map[string]model.NewBaseModel = model.GetRegistryAllConstructModelBase()
			for _, constructor := range modelsConstructors {
				modelsToMigrate = append(modelsToMigrate, constructor())
			}
			return modelsToMigrate
		}(),
		dbManager: func() *db.Manager {
			var dbManager *db.Manager = container.Inject[_db.Manager]().(*db.Manager)
			return dbManager
		}(),
	}
}

func (m *MigrateAllModels) Execute() error {
	var err error = m.dbManager.ExecuteInTransaction(func(ctx _db.Context) error {
		var dbCtx *db.Context = ctx.(*db.Context)
		if err := dbCtx.
			DB().
			AutoMigrate(m.models...); err != nil {
			return err
		}
		return nil
	})
	return err
}
