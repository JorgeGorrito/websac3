package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerCreateAccessRequest(routerGroup *gin.RouterGroup) {
	createAccessRequestController := controller.InitCreateAccessRequestController()

	routerGroup.POST("/access-request", createAccessRequestController.CreateAccessRequest)
}

func RegisterAccessRequest(routerGroup *gin.RouterGroup) {
	registerCreateAccessRequest(routerGroup)
}
