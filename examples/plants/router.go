package plants

import (
	"darling"
)

var Routers *darling.ControllerRegistor

func init() {
	Routers = darling.NewControllerRegistor()
	Routers.Add("/v1/plants/(\\w+)", &PlantCtrl{})
	Routers.Add("/v1/plants", &PlantListCtrl{})
}
