package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerCreateCourseRoute(routerGroup *gin.RouterGroup) {
	var createCourseController = controller.GetCreateCourseController()
	routerGroup.POST("/course", createCourseController.CreateCourse)
}

func registerListCourseByDegreeProgramRoute(routerGroup *gin.RouterGroup) {
	var listCourseByDegreeProgramController = controller.GetListCourseByDegreeProgramController()
	routerGroup.GET("/degree-program/:degree_program_id/courses", listCourseByDegreeProgramController.Handle)
}

func registerGetCourseByIDRoute(routerGroup *gin.RouterGroup) {
	var getCourseByIDController = controller.GetGetCourseByIDController()
	routerGroup.GET("/course/:course_id", getCourseByIDController.Handle)
}

func registerGetCourseTopicsByCourseIDRoute(routerGroup *gin.RouterGroup) {
	var getCourseTopicsByCourseIDController = controller.GetGetCourseTopicsByCourseIDController()
	routerGroup.GET("/course/:course_id/topics", getCourseTopicsByCourseIDController.Handle)
}

func RegisterCourseRoutes(routerGroup *gin.RouterGroup) {
	authGroup := getAuthRequiredGroup(routerGroup)
	registerCreateCourseRoute(authGroup)
	registerListCourseByDegreeProgramRoute(authGroup)
	registerGetCourseByIDRoute(authGroup)
	registerGetCourseTopicsByCourseIDRoute(authGroup)
}
