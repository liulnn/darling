package plants

import (
	"darling"
)

var Routers map[string]darling.ControllerInterface

func init() {
	Routers = make(map[string]darling.ControllerInterface)
	Routers["/v1/plants/(\\w+)"] = &PlantCtrl{}
	Routers["/v1/plants"] = &PlantListCtrl{}
}
