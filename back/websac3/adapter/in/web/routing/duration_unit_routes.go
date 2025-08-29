package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerListDurationUnitRoute(routerGroup *gin.RouterGroup) {
	var listDurationUnitController = controller.GetListDurationUnitController()
	routerGroup.GET("/duration-unit", listDurationUnitController.Handle)
}

func RegisterDurationUnitRoutes(routerGroup *gin.RouterGroup) {
	authGroup := getAuthRequiredGroup(routerGroup)
	registerListDurationUnitRoute(authGroup)
}
