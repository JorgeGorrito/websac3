package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerCreateDegreeProgramRoute(routerGroup *gin.RouterGroup) {
	var createDegreeProgramController = controller.GetCreateDegreeProgramController()
	routerGroup.POST("/degree-program", createDegreeProgramController.CreateDegreeProgram)
}

func RegisterDegreeProgramRoutes(routerGroup *gin.RouterGroup) {
	authGroup := getAuthRequiredGroup(routerGroup)
	registerCreateDegreeProgramRoute(authGroup)
}
