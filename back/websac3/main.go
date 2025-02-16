package main

import (
	"websac3/config"
	"websac3/dependencies"
	"websac3/routing"

	"github.com/JorgeGorrito/anise-with-gin/anise"
	"github.com/gin-gonic/gin"
)

func main() {
	if app, err := anise.NewWebApplication(
		gin.Default(),
		config.GetManager(),
		routing.GetManager(),
		dependencies.GetManager(),
	); err != nil {
		panic(err)
	} else {
		app.Run()
	}
}
