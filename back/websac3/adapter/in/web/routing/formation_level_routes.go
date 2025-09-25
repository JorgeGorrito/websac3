package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerListFormationLevelRoute(routerGroup *gin.RouterGroup) {
	var listFormationLevelController = controller.GetListFormationLevelController()
	routerGroup.GET("/formation-level", listFormationLevelController.Handle)
}

func RegisterFormationLevelRoutes(routerGroup *gin.RouterGroup) {
	authGroup := getAuthRequiredGroup(routerGroup)
	registerListFormationLevelRoute(authGroup)
}
