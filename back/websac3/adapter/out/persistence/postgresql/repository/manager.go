package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type manager struct {
	db *gorm.DB
}

var instance *manager = nil

func initManager() (err error) {
	requiredEnv := []string{"DB_HOST", "DB_USER", "DB_NAME"}
	for _, env := range requiredEnv {
		if os.Getenv(env) == "" {
			return fmt.Errorf("variable de entorno requerida faltante: %s", env)
		}
	}

	config := struct {
		host     string
		user     string
		password string
		dbname   string
		port     string
		sslmode  string
		timezone string
	}{
		host:     os.Getenv("DB_HOST"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
		sslmode:  os.Getenv("DB_SSLMODE"),
		timezone: os.Getenv("DB_TIMEZONE"),
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.host,
		config.user,
		config.password,
		config.dbname,
		config.port,
		config.sslmode,
		config.timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error al conectar a la base de datos: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error al obtener instancia DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	instance = &manager{db: db}
	return nil
}

func GetManager() (*manager, error) {
	if instance == nil {
		err := initManager()
		if err != nil {
			return nil, err
		}
	}
	return instance, nil
}
