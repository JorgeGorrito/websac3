package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerListHigherEducationInstitutionRoutes(routerGroup *gin.RouterGroup) {
	var listHigherEducationInstitutionController = controller.GetListHigherEducationInstitutionController()
	routerGroup.GET("/higher-education-institution", listHigherEducationInstitutionController.Handle)
}

func RegisterHigherEducationInstitutionRoutes(routerGroup *gin.RouterGroup) {
	registerListHigherEducationInstitutionRoutes(routerGroup)
}
