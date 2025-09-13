package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerListRoles(routerGroup *gin.RouterGroup) {
	listRoleController := controller.GetListRoleController()

	routerGroup.GET("/roles", listRoleController.Handle)
}

func RegisterRoleRoutes(routerGroup *gin.RouterGroup) {
	registerListRoles(routerGroup)
}
