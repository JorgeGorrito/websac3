package command

import (
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-with-gin/anise/command"
	"gorm.io/gorm"
)

type DBReset struct {
	cmdprinter command.CMDPrinter
	db         *gorm.DB
	models     []any
	dbManager  *db.Manager
}

func NewDBReset(params map[string]string, cmdprinter command.CMDPrinter) command.Command {
	return &DBReset{
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

func (m *DBReset) Execute() error {
	return m.dbManager.ExecuteInTransaction(func(ctx _db.Context) error {
		var dbCtx *db.Context = ctx.(*db.Context)

		// 1. Deshabilitar foreign key checks temporalmente
		if err := dbCtx.DB().Exec("SET session_replication_role = replica;").Error; err != nil {
			return err
		}

		// 2. Eliminar y recrear el schema completo (más robusto que eliminar tablas individuales)
		if err := dbCtx.DB().Exec("DROP SCHEMA IF EXISTS public CASCADE").Error; err != nil {
			return err
		}

		if err := dbCtx.DB().Exec("CREATE SCHEMA public").Error; err != nil {
			return err
		}

		// 3. Restaurar permisos del schema
		if err := dbCtx.DB().Exec("GRANT ALL ON SCHEMA public TO public").Error; err != nil {
			return err
		}

		// 4. Rehabilitar foreign key checks
		if err := dbCtx.DB().Exec("SET session_replication_role = DEFAULT;").Error; err != nil {
			return err
		}

		// 5. Ejecutar migraciones para recrear todas las tablas
		if err := dbCtx.DB().AutoMigrate(m.models...); err != nil {
			return err
		}

		// 6. Ejecutar AfterMigrate para cada modelo (para crear índices personalizados)
		for _, modelInstance := range m.models {
			if migrator, ok := modelInstance.(interface{ AfterMigrate(*gorm.DB) error }); ok {
				if err := migrator.AfterMigrate(dbCtx.DB()); err != nil {
					return err
				}
			}
		}

		return nil
	})
}
