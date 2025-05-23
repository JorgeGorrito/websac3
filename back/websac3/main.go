package main

import (
	"websac3/common/dependencies"
	"websac3/common/dependencies/container"

	"github.com/JorgeGorrito/anise-with-gin/anise"
	"github.com/JorgeGorrito/anise-with-gin/anise/command"
	"github.com/JorgeGorrito/anise-with-gin/anise/config"
	"github.com/JorgeGorrito/anise-with-gin/anise/routing"
	"github.com/gin-gonic/gin"
)

func main() {
	dependencies.InitDependenciesManager()
	app := anise.NewWebApplication(
		gin.Default(),
		container.Inject[config.Manager](),
		container.Inject[routing.Manager](),
		container.Inject[command.Manager](),
	)
	app.Run(":8110")
}
