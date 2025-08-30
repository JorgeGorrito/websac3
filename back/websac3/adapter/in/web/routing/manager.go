package routing

import (
	"os"
	"websac3/adapter/in/web/middleware"

	"github.com/gin-gonic/gin"
)

type manager struct{}

func NewRoutingManager() *manager {
	return &manager{}
}

func getAuthRequiredGroup(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	var secretKey = os.Getenv("JWT_SECRET_KEY")
	authRequired := routerGroup.Group("")
	authRequired.Use(middleware.AuthMiddleware([]byte(secretKey)))
	return authRequired
}

func (m *manager) RegisterRoutes(engine *gin.Engine) error {
	m.configureCors(engine)

	var routerGroup *gin.RouterGroup = engine.Group("/api/v1/:lang")
	RegisterAccessRequestRoutes(routerGroup)
	RegisterIdentificationTypesRoutes(routerGroup)
	RegisterHigherEducationInstitutionRoutes(routerGroup)
	RegisterTopicRoutes(routerGroup)
	RegisterDurationUnitRoutes(routerGroup)
	RegisterDegreeProgramRoutes(routerGroup)
	RegisterAuthRoutes(routerGroup)
	RegisterSwagger(engine)

	return nil
}

func (m *manager) configureCors(engine *gin.Engine) {
	engine.Use(middleware.CorsMiddleware())
}

var instance *manager = nil

func GetManager() *manager {
	if instance == nil {
		instance = &manager{}
	}
	return instance
}
