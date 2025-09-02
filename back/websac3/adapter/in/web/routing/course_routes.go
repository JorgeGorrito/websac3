package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerCreateCourseRoute(routerGroup *gin.RouterGroup) {
	var createCourseController = controller.GetCreateCourseController()
	routerGroup.POST("/course", createCourseController.CreateCourse)
}

func RegisterCourseRoutes(routerGroup *gin.RouterGroup) {
	authGroup := getAuthRequiredGroup(routerGroup)
	registerCreateCourseRoute(authGroup)
}
