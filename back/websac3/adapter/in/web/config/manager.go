package config

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type manager struct{}

func NewConfigManager() *manager {
	return &manager{}
}

func (m *manager) loadEnviroment(errorList *error) {
	if err := godotenv.Load(); err != nil {
		*errorList = errors.Join(*errorList, err)
	}
}

func (m *manager) ConfigureApplication() error {
	var errorList error
	m.loadEnviroment(&errorList)
	return errorList
}

func (m *manager) ConfigureEngine(engine *gin.Engine) error {
	var errorList error

	return errorList
}
