package main

import (
	"darling"
	"plant-server/plants"
)

func routerConfig(app *darling.App) {
	app.Handlers.Add("/v1/plants/(\\w+)", &plants.PlantCtrl{})
	app.Handlers.Add("/v1/plants", &plants.PlantListCtrl{})
}

func main() {
	var app = darling.NewApp()
	routerConfig(app)
	app.Run("127.0.0.1", 7432)
}
