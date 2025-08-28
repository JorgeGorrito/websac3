package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerListTopicRoute(routerGroup *gin.RouterGroup) {
	var listTopicController = controller.GetListTopicController()
	routerGroup.GET("/topic", listTopicController.Handle)
}

func RegisterTopicRoutes(routerGroup *gin.RouterGroup) {
	authGroup := getAuthRequiredGroup(routerGroup)
	registerListTopicRoute(authGroup)
}
