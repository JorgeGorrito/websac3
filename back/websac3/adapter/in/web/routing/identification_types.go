package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerListIdentificationTypes(routerGroup *gin.RouterGroup) {
	listIdentificationTypeController := controller.GetListIdentificationTypeController()

	routerGroup.GET("/identification-types", listIdentificationTypeController.Handle)
}

func RegisterIdentificationTypesRoutes(routerGroup *gin.RouterGroup) {
	registerListIdentificationTypes(routerGroup)
}
