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

func registerEvaluateDegreeProgramRoute(routerGroup *gin.RouterGroup) {
	var evaluateDegreeProgramController = controller.GetEvaluateDegreeProgramController()
	routerGroup.POST("/degree-program/evaluate", evaluateDegreeProgramController.Handle)
}

func registerListReportsByDegreeProgramRoute(routerGroup *gin.RouterGroup) {
	var listReportsByDegreeProgramController = controller.GetListReportsByDegreeProgramController()
	routerGroup.GET("/degree-program/:degree_program_id/reports", listReportsByDegreeProgramController.Handle)
}

func registerGetDegreeProgramByIDRoute(routerGroup *gin.RouterGroup) {
	var getDegreeProgramByIDController = controller.GetGetDegreeProgramByIDController()
	routerGroup.GET("/degree-program/:degree_program_id", getDegreeProgramByIDController.Handle)
}

func registerUpdateDegreeProgramRoute(routerGroup *gin.RouterGroup) {
	var updateDegreeProgramController = controller.GetUpdateDegreeProgramController()
	routerGroup.PUT("/degree-program/:degree_program_id", updateDegreeProgramController.UpdateDegreeProgram)
}

func RegisterDegreeProgramRoutes(routerGroup *gin.RouterGroup) {
	authGroup := getAuthRequiredGroup(routerGroup)
	registerCreateDegreeProgramRoute(authGroup)
	registerListDegreeProgramRoute(authGroup)
	registerEvaluateDegreeProgramRoute(authGroup)
	registerListReportsByDegreeProgramRoute(authGroup)
	registerUpdateDegreeProgramRoute(authGroup)
	registerGetDegreeProgramByIDRoute(routerGroup) // Sin autenticación requerida
}
