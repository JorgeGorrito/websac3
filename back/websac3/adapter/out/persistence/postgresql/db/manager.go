package db

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {
	var errorList error
	var err error
	var conn *gorm.DB = nil

	requiredEnv := []string{"DB_HOST", "DB_USER", "DB_NAME"}
	for _, env := range requiredEnv {
		if os.Getenv(env) == "" {
			return nil, fmt.Errorf("variable de entorno requerida faltante: %s", env)
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

	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		errorList = errors.Join(errorList, err)
	}

	sqlDB, err := conn.DB()
	if err != nil {
		errorList = errors.Join(errorList, err)
	}

	maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		errorList = errors.Join(errorList, err)
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		errorList = errors.Join(errorList, err)
	}

	connMaxLifetime, err := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME"))
	if err != nil {
		errorList = errors.Join(errorList, err)
	}

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	return conn, errorList
}
