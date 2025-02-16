package routing

import (
	"github.com/JorgeGorrito/anise-with-gin/anise/dependencies"
	"github.com/gin-gonic/gin"
)

type manager struct{}

func (m *manager) RegisterRoutes(engine *gin.Engine, dependenciesResolver dependencies.Resolver) error {
	return nil
}

var instance *manager = nil

func GetManager() *manager {
	if instance == nil {
		instance = &manager{}
	}
	return instance
}
