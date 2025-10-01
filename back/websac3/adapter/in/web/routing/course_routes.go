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

func registerListCourseTypesRoute(routerGroup *gin.RouterGroup) {
	var listCourseTypesController = controller.GetListCourseTypesController()
	routerGroup.GET("/course-types", listCourseTypesController.Handle)
}

func registerListCourseNaturesRoute(routerGroup *gin.RouterGroup) {
	var listCourseNaturesController = controller.GetListCourseNaturesController()
	routerGroup.GET("/course-natures", listCourseNaturesController.Handle)
}

func registerDeleteCourseRoute(routerGroup *gin.RouterGroup) {
	var deleteCourseController = controller.GetDeleteCourseController()
	routerGroup.DELETE("/course/:course_id", deleteCourseController.Handle)
}

func registerUpdateCourseRoute(routerGroup *gin.RouterGroup) {
	var updateCourseController = controller.GetUpdateCourseController()
	routerGroup.PUT("/course/:course_id", updateCourseController.Handle)
}

func registerDownloadCourseTemplateRoute(routerGroup *gin.RouterGroup) {
	var downloadTemplateController = controller.GetDownloadCourseTemplateController()
	routerGroup.GET("/course/template", downloadTemplateController.DownloadTemplate)
}

func registerBulkCreateCourseRoute(routerGroup *gin.RouterGroup) {
	var bulkCreateController = controller.GetBulkCreateCourseController()
	routerGroup.POST("/course/bulk", bulkCreateController.BulkCreateCourse)
}

func RegisterCourseRoutes(routerGroup *gin.RouterGroup) {
	authGroup := getAuthRequiredGroup(routerGroup)
	registerCreateCourseRoute(authGroup)
	registerListCourseByDegreeProgramRoute(authGroup)
	registerGetCourseByIDRoute(authGroup)
	registerGetCourseTopicsByCourseIDRoute(authGroup)
	registerUpdateCourseRoute(authGroup)
	registerDeleteCourseRoute(authGroup)
	registerListCourseTypesRoute(authGroup)
	registerListCourseNaturesRoute(authGroup)
	registerDownloadCourseTemplateRoute(authGroup)
	registerBulkCreateCourseRoute(authGroup)
}
