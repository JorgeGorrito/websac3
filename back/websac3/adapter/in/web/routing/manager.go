package routing

import (
	"github.com/gin-gonic/gin"
)

type manager struct{}

func NewRoutingManager() *manager {
	return &manager{}
}

func (m *manager) RegisterRoutes(engine *gin.Engine) error {
	var routerGroup *gin.RouterGroup = engine.Group("/api/v1")
	RegisterAccessRequest(routerGroup)

	return nil
}

var instance *manager = nil

func GetManager() *manager {
	if instance == nil {
		instance = &manager{}
	}
	return instance
}
