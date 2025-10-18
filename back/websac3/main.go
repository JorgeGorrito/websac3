package main

import (
	"websac3/common/dependencies"
	"websac3/common/dependencies/container"

	_ "websac3/docs"

	"github.com/JorgeGorrito/anise-with-gin/anise"
	"github.com/JorgeGorrito/anise-with-gin/anise/command"
	"github.com/JorgeGorrito/anise-with-gin/anise/config"
	"github.com/JorgeGorrito/anise-with-gin/anise/routing"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Websac3 API
// @version 1.0
// @description API documentation for Websac3 backend
// @host localhost:8110
// @schemes https
// @contact.name API Support
// @contact.email j0rg3.4b3ll4@gmail.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file. \n Error: " + err.Error())
	}
	dependencies.InitDependenciesManager()
	app := anise.NewWebApplication()
	app.
		SetEngine(gin.Default()).
		SetConfigManager(container.Inject[config.Manager]()).
		SetRoutesManager(container.Inject[routing.Manager]()).
		SetCommandsManager(container.Inject[command.Manager]()).
		RunTLS(":8110", "./localhost.crt", "./localhost.key")
}
