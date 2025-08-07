package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerCreateAccessRequest(routerGroup *gin.RouterGroup) {
	createAccessRequestController := controller.GetCreateAccessRequestController()

	routerGroup.POST("/access-request", createAccessRequestController.CreateAccessRequest)
}

func registerValidateEmailRequest(routerGroup *gin.RouterGroup) {
	validateEmailController := controller.GetValidateEmailController()

	routerGroup.POST("/access-request/email/validate", validateEmailController.ValidateEmail)
}

func registerListAccessRequest(routerGroup *gin.RouterGroup) {
	listAccessRequestController := controller.GetListAccessRequestController()

	routerGroup.GET("/access-request", listAccessRequestController.Handle)
}

func RegisterAccessRequestRoutes(routerGroup *gin.RouterGroup) {
	authRequiredGroup := getAuthRequiredGroup(routerGroup)

	registerCreateAccessRequest(routerGroup)
	registerValidateEmailRequest(routerGroup)
	registerListAccessRequest(authRequiredGroup)
}
