package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerListProfessionalRoles(routerGroup *gin.RouterGroup) {
	listProfessionalRoleController := controller.GetListProfessionalRoleController()

	routerGroup.GET("/professional-roles", listProfessionalRoleController.Handle)
}

func RegisterProfessionalRoleRoutes(routerGroup *gin.RouterGroup) {
	registerListProfessionalRoles(routerGroup)
}
