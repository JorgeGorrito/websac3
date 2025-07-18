package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerCreateAccessRequest(routerGroup *gin.RouterGroup) {
	createAccessRequestController := controller.GetCreateAccessRequestController()

	routerGroup.POST("/access-request", createAccessRequestController.CreateAccessRequest)
}

func validateEmailRequest(routerGroup *gin.RouterGroup) {
	validateEmailController := controller.GetValidateEmailController()

	routerGroup.POST("/access-request/email/validate", validateEmailController.ValidateEmail)
}

func RegisterAccessRequest(routerGroup *gin.RouterGroup) {
	registerCreateAccessRequest(routerGroup)
	validateEmailRequest(routerGroup)
}
