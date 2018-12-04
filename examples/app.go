package main

import (
	"darling"
	"./plants"
)

func routerConfig(app *darling.App) {
	app.Handlers.Extend(plants.Routers)
}

func main() {
	var app = darling.NewApp()
	routerConfig(app)
	app.Run("127.0.0.1", 7432)
}
