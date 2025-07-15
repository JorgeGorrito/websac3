package routing

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwagger(engine *gin.Engine) error {
	engine.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
	return nil
}
