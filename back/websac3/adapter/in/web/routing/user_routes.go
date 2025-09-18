package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerListUsersRoute(routerGroup *gin.RouterGroup) {
	listUsersController := controller.GetListUsersController()
	routerGroup.GET("", listUsersController.Handle)
}

func registerDeactivateUserRoute(routerGroup *gin.RouterGroup) {
	deactivateUserController := controller.GetDeactivateUserController()
	routerGroup.PUT("/:user_id/deactivate", deactivateUserController.Handle)
}

func registerActivateUserRoute(routerGroup *gin.RouterGroup) {
	activateUserController := controller.GetActivateUserController()
	routerGroup.PUT("/:user_id/activate", activateUserController.Handle)
}

func RegisterUserRoutes(routerGroup *gin.RouterGroup) {
	userGroup := routerGroup.Group("/users")
	authGroup := getAuthRequiredGroup(userGroup)
	registerListUsersRoute(authGroup)
	registerDeactivateUserRoute(authGroup)
	registerActivateUserRoute(authGroup)
}
