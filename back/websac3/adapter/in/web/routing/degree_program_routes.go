package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerCreateDegreeProgramRoute(routerGroup *gin.RouterGroup) {
	var createDegreeProgramController = controller.GetCreateDegreeProgramController()
	routerGroup.POST("/degree-program", createDegreeProgramController.CreateDegreeProgram)
}

func registerListDegreeProgramRoute(routerGroup *gin.RouterGroup) {
	var listDegreeProgramController = controller.GetListDegreeProgramController()
	routerGroup.GET("/degree-program", listDegreeProgramController.Handle)
}

func RegisterDegreeProgramRoutes(routerGroup *gin.RouterGroup) {
	authGroup := getAuthRequiredGroup(routerGroup)
	registerCreateDegreeProgramRoute(authGroup)
	registerListDegreeProgramRoute(authGroup)
}
