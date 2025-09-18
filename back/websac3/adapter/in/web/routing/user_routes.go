package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerListUsersRoute(routerGroup *gin.RouterGroup) {
	listUsersController := controller.GetListUsersController()
	routerGroup.GET("", listUsersController.Handle)
}

func RegisterUserRoutes(routerGroup *gin.RouterGroup) {
	userGroup := routerGroup.Group("/users")
	authGroup := getAuthRequiredGroup(userGroup)
	registerListUsersRoute(authGroup)
}
