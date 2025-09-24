package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerCreateExpertConsultation(routerGroup *gin.RouterGroup) {
	createExpertConsultationController := controller.GetCreateExpertConsultationController()

	routerGroup.POST("/expert-consultation", createExpertConsultationController.CreateExpertConsultation)
}

func RegisterExpertConsultationRoutes(routerGroup *gin.RouterGroup) {
	authRequiredGroup := getAuthRequiredGroup(routerGroup)

	registerCreateExpertConsultation(authRequiredGroup)
}
