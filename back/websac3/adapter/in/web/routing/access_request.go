package routing

import (
	"websac3/adapter/in/web/controller"
	"websac3/app/port/in/usecase"
	"websac3/common/dependencies/container"

	"github.com/gin-gonic/gin"
)

func registerCreateAccessRequest(routerGroup *gin.RouterGroup) {
	createAccessRequestController := controller.InitCreateAccessRequestController(
		container.Inject[usecase.CreateAccessRequestUseCase](),
	)

	routerGroup.POST("/access-request", createAccessRequestController.CreateAccessRequest)
}

func RegisterAccessRequest(routerGroup *gin.RouterGroup) {
	registerCreateAccessRequest(routerGroup)
}
