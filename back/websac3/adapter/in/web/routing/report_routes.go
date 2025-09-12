package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerGetReportByIDRoute(routerGroup *gin.RouterGroup) {
	var getReportByIDController = controller.GetGetReportByIDController()
	routerGroup.GET("/report/:report_id", getReportByIDController.Handle)
}

func RegisterReportRoutes(routerGroup *gin.RouterGroup) {
	authGroup := getAuthRequiredGroup(routerGroup)
	registerGetReportByIDRoute(authGroup)
}
