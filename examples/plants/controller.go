package plants

import (
	"darling"
	"fmt"
)

type PlantListCtrl struct {
	darling.Controller
}

func (c *PlantListCtrl) Post() {
	req := c.Request.In
	req.ParseForm()
	name := req.PostFormValue("name")
	fmt.Println("name:", name)

	c.Response.StatusCode = 201
}

func (c *PlantListCtrl) Get() {

	c.Response.StatusCode = 200
}

type PlantCtrl struct {
	darling.Controller
	plantId string
}

func (c *PlantCtrl) Prepare() {
	plantId := c.PathParams[0]
	c.plantId = plantId
}

func (c *PlantCtrl) Get() {
	fmt.Println("plantId:", c.plantId)
	c.Response.StatusCode = 200
}

func (c *PlantCtrl) Put() {
	fmt.Println("plantId:", c.plantId)
	c.Response.StatusCode = 200
}
