package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerCreateReportFeedbackRoute(routerGroup *gin.RouterGroup) {
	createReportFeedbackController := controller.GetCreateReportFeedbackController()
	routerGroup.POST("/report-feedbacks", createReportFeedbackController.Handle)
}

func registerListReportsPendingFeedbackRoute(routerGroup *gin.RouterGroup) {
	listReportsPendingFeedbackController := controller.NewListReportsPendingFeedbackController()
	routerGroup.GET("/report-feedbacks/pending-reports", listReportsPendingFeedbackController.Handle)
}

func registerGetReportFeedbackRoute(routerGroup *gin.RouterGroup) {
	getReportFeedbackController := controller.NewGetReportFeedbackController()
	routerGroup.GET("/report-feedbacks/report/:reportId", getReportFeedbackController.Handle)
}

func registerListReportFeedbacksRoute(routerGroup *gin.RouterGroup) {
	listReportFeedbacksController := controller.NewListReportFeedbacksController()
	routerGroup.GET("/report-feedbacks", listReportFeedbacksController.Handle)
}

func RegisterReportFeedbackRoutes(routerGroup *gin.RouterGroup) {
	authGroup := getAuthRequiredGroup(routerGroup)
	registerCreateReportFeedbackRoute(authGroup)
	registerListReportFeedbacksRoute(authGroup)
	registerListReportsPendingFeedbackRoute(authGroup)
	registerGetReportFeedbackRoute(authGroup)
}
