package darling

import (
	"net/http"
)

type Controller struct {
	Request  *http.Request
	Response http.ResponseWriter
	Name     string
}

type ControllerInterface interface {
	Init(req *http.Request, resp http.ResponseWriter, name string)
	Prepare()
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	Finish()
}

func (c *Controller) Init(req *http.Request, resp http.ResponseWriter, name string) {
	c.Request = req
	c.Response = resp
	c.Name = name
}

func (c *Controller) Prepare() {

}

func (c *Controller) Finish() {

}

func (c *Controller) Get() {
	http.Error(c.Response, "Method Not Allowed", 405)
}

func (c *Controller) Post() {
	http.Error(c.Response, "Method Not Allowed", 405)
}

func (c *Controller) Delete() {
	http.Error(c.Response, "Method Not Allowed", 405)
}

func (c *Controller) Put() {
	http.Error(c.Response, "Method Not Allowed", 405)
}

func (c *Controller) Head() {
	http.Error(c.Response, "Method Not Allowed", 405)
}

func (c *Controller) Patch() {
	http.Error(c.Response, "Method Not Allowed", 405)
}

func (c *Controller) Options() {
	http.Error(c.Response, "Method Not Allowed", 405)
}
