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

func registerCreateUserFromToken(routerGroup *gin.RouterGroup) {
	createUserFromTokenController := controller.GetCreateUserFromTokenController()

	routerGroup.POST("/user/create-from-token", createUserFromTokenController.CreateUserFromToken)
}

func registerListAccessRequest(routerGroup *gin.RouterGroup) {
	listAccessRequestController := controller.GetListAccessRequestController()

	routerGroup.GET("/access-request", listAccessRequestController.Handle)
}

func registerApproveAccessRequest(routerGroup *gin.RouterGroup) {
	approveAccessRequestController := controller.GetApproveAccessRequestController()

	routerGroup.POST("/access-request/:accessRequestID/approve", approveAccessRequestController.Handle)
}

func registerRejectAccessRequest(routerGroup *gin.RouterGroup) {
	rejectAccessRequestController := controller.GetRejectAccessRequestController()

	routerGroup.POST("/access-request/:accessRequestID/reject", rejectAccessRequestController.Handle)
}

func RegisterAccessRequestRoutes(routerGroup *gin.RouterGroup) {
	authRequiredGroup := getAuthRequiredGroup(routerGroup)

	registerCreateAccessRequest(routerGroup)
	registerValidateEmailRequest(routerGroup)
	registerCreateUserFromToken(routerGroup)

	registerListAccessRequest(authRequiredGroup)
	registerApproveAccessRequest(authRequiredGroup)
	registerRejectAccessRequest(authRequiredGroup)
}
