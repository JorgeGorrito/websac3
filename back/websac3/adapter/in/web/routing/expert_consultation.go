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

func registerListAnsweredExpertConsultations(routerGroup *gin.RouterGroup) {
	listAnsweredExpertConsultationsController := controller.NewListAnsweredExpertConsultationsController()

	routerGroup.GET("/expert-consultations/answered", listAnsweredExpertConsultationsController.Handle)
}

func registerAcceptExpertConsultation(routerGroup *gin.RouterGroup) {
	acceptExpertConsultationController := controller.NewAcceptExpertConsultationController()

	routerGroup.PUT("/expert-consultations/:consultation_id/accept", acceptExpertConsultationController.Handle)
}

func registerRejectExpertConsultation(routerGroup *gin.RouterGroup) {
	rejectExpertConsultationController := controller.NewRejectExpertConsultationController()

	routerGroup.PUT("/expert-consultations/:consultation_id/reject", rejectExpertConsultationController.Handle)
}

func registerGetExpertConsultationByID(routerGroup *gin.RouterGroup) {
	getExpertConsultationByIDController := controller.NewGetExpertConsultationByIDController()

	routerGroup.GET("/expert-consultations/:consultation_id", getExpertConsultationByIDController.Handle)
}

func registerListExpertConsultationStatuses(routerGroup *gin.RouterGroup) {
	listExpertConsultationStatusController := controller.GetListExpertConsultationStatusController()

	routerGroup.GET("/expert-consultation-statuses", listExpertConsultationStatusController.Handle)
}

func RegisterExpertConsultationRoutes(routerGroup *gin.RouterGroup) {
	authRequiredGroup := getAuthRequiredGroup(routerGroup)

	registerCreateExpertConsultation(authRequiredGroup)
	registerListUserExpertConsultations(authRequiredGroup)
	registerListPendingExpertConsultations(authRequiredGroup)
	registerListAnsweredExpertConsultations(authRequiredGroup)
	registerGetExpertConsultationByID(authRequiredGroup)
	registerAcceptExpertConsultation(authRequiredGroup)
	registerRejectExpertConsultation(authRequiredGroup)
	registerListExpertConsultationStatuses(routerGroup) // Sin autenticación requerida
}
