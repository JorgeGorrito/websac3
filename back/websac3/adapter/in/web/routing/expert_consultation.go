package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerCreateExpertConsultation(routerGroup *gin.RouterGroup) {
	createExpertConsultationController := controller.GetCreateExpertConsultationController()

	routerGroup.POST("/expert-consultation", createExpertConsultationController.CreateExpertConsultation)
}

func registerListUserExpertConsultations(routerGroup *gin.RouterGroup) {
	listUserExpertConsultationsController := controller.NewListUserExpertConsultationsController()

	routerGroup.GET("/user/expert-consultations", listUserExpertConsultationsController.Handle)
}

func registerListPendingExpertConsultations(routerGroup *gin.RouterGroup) {
	listPendingExpertConsultationsController := controller.NewListPendingExpertConsultationsController()

	routerGroup.GET("/expert-consultations/pending", listPendingExpertConsultationsController.Handle)
}

func registerAcceptExpertConsultation(routerGroup *gin.RouterGroup) {
	acceptExpertConsultationController := controller.NewAcceptExpertConsultationController()

	routerGroup.PUT("/expert-consultations/:consultation_id/accept", acceptExpertConsultationController.Handle)
}

func registerRejectExpertConsultation(routerGroup *gin.RouterGroup) {
	rejectExpertConsultationController := controller.NewRejectExpertConsultationController()

	routerGroup.PUT("/expert-consultations/:consultation_id/reject", rejectExpertConsultationController.Handle)
}

func RegisterExpertConsultationRoutes(routerGroup *gin.RouterGroup) {
	authRequiredGroup := getAuthRequiredGroup(routerGroup)

	registerCreateExpertConsultation(authRequiredGroup)
	registerListUserExpertConsultations(authRequiredGroup)
	registerListPendingExpertConsultations(authRequiredGroup)
	registerAcceptExpertConsultation(authRequiredGroup)
	registerRejectExpertConsultation(authRequiredGroup)
}
