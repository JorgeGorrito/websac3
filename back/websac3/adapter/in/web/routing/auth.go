package routing

import (
	"websac3/adapter/in/web/controller"

	"github.com/gin-gonic/gin"
)

func registerLoginRoute(routerGroup *gin.RouterGroup) {
	loginController := controller.GetLoginController()
	routerGroup.POST("", loginController.Handle)
}

func registerRefreshTokenRoute(routerGroup *gin.RouterGroup) {
	refreshTokenController := controller.GetRefreshTokenController()
	routerGroup.POST("/refresh", refreshTokenController.Handle)
}

func RegisterAuthRoutes(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/auth")
	registerLoginRoute(authGroup)
	registerRefreshTokenRoute(authGroup)
}
