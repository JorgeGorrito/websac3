package config

import "github.com/gin-gonic/gin"

type manager struct{}

func (m *manager) ConfigureEngine(engine *gin.Engine) error {
	return nil
}

var instance *manager = nil

func GetManager() *manager {
	if instance == nil {
		instance = &manager{}
	}
	return instance
}
