package command

import (
	"fmt"
	"time"
	"websac3/adapter/out/persistence/postgresql/db"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/adapter/out/persistence/postgresql/seeders"
	_db "websac3/app/port/out/persistence/db"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-with-gin/anise/command"
	"gorm.io/gorm"
)

type InitApplicationCommand struct {
	cmdPrinter command.CMDPrinter
	dbManager  *db.Manager
	models     []any
}

func NewInitApplicationCommand(params map[string]string, cmdPrinter command.CMDPrinter) command.Command {
	return &InitApplicationCommand{
		cmdPrinter: cmdPrinter,
		dbManager: func() *db.Manager {
			var dbManager *db.Manager = container.Inject[_db.Manager]().(*db.Manager)
			return dbManager
		}(),
		models: func() (modelsToMigrate []any) {
			var modelsConstructors map[string]model.NewBaseModel = model.GetRegistryAllConstructModelBase()
			for _, constructor := range modelsConstructors {
				modelsToMigrate = append(modelsToMigrate, constructor())
			}
			return modelsToMigrate
		}(),
	}
}

func (c *InitApplicationCommand) Execute() error {
	return c.dbManager.ExecuteInTransaction(func(ctx _db.Context) error {
		var dbCtx *db.Context = ctx.(*db.Context)
		var database *gorm.DB = dbCtx.DB()

		// Primero, migrar solo la tabla de migraciones
		fmt.Println("Verificando tabla de migraciones...")
		if err := database.AutoMigrate(&model.Migration{}); err != nil {
			return fmt.Errorf("error al migrar tabla de migraciones: %w", err)
		}

		// Verificar si ya existe la migración inicial
		var migration model.Migration
		result := database.Where("name = ?", "initial_migration").First(&migration)

		if result.Error == nil {
			// La migración ya existe
			fmt.Println("✓ La aplicación ya ha sido inicializada previamente.")
			fmt.Printf("  Migración ejecutada el: %s\n", migration.ExecutedAt.Format("2006-01-02 15:04:05"))
			return nil
		}

		if result.Error != gorm.ErrRecordNotFound {
			// Error diferente a "registro no encontrado"
			return fmt.Errorf("error al verificar migraciones: %w", result.Error)
		}

		// No existe la migración, proceder con la inicialización
		fmt.Println("Iniciando aplicación por primera vez...")

		// Migrar todos los modelos
		fmt.Println("Migrando modelos...")
		if err := database.AutoMigrate(c.models...); err != nil {
			return fmt.Errorf("error al migrar modelos: %w", err)
		}
		fmt.Println("✓ Modelos migrados exitosamente")

		// Ejecutar seeders esenciales
		fmt.Println("Ejecutando seeders esenciales...")
		var essentialDataSeeder seeders.Seeder = seeders.EssentialData()
		if err := essentialDataSeeder.Seed(ctx); err != nil {
			return fmt.Errorf("error al ejecutar seeders esenciales: %w", err)
		}
		fmt.Println("✓ Seeders esenciales ejecutados exitosamente")

		// Ejecutar seeder de roles profesionales
		fmt.Println("Ejecutando seeder de roles profesionales...")
		var professionalRolesSeeder seeders.Seeder = seeders.ProfessionalRoles()
		if err := professionalRolesSeeder.Seed(ctx); err != nil {
			return fmt.Errorf("error al ejecutar seeder de roles profesionales: %w", err)
		}
		fmt.Println("✓ Roles profesionales ejecutados exitosamente")

		// Registrar la migración como completada
		migration = model.Migration{
			Name:        "initial_migration",
			ExecutedAt:  time.Now(),
			Description: "Migración inicial de la base de datos, seeders esenciales y roles profesionales",
		}
		if err := database.Create(&migration).Error; err != nil {
			return fmt.Errorf("error al registrar migración: %w", err)
		}

		fmt.Println("✓ Aplicación inicializada correctamente")
		return nil
	})
}
