package main

import (
	"websac3/common/dependencies"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-with-gin/anise"
	"github.com/JorgeGorrito/anise-with-gin/anise/command"
	"github.com/JorgeGorrito/anise-with-gin/anise/config"
	"github.com/JorgeGorrito/anise-with-gin/anise/routing"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
		Run(":8110")
}
