package darling

import (
	"net/http"
)

type Controller struct {
	Request    *http.Request
	Response   http.ResponseWriter
	PathParams []string
}

type ControllerInterface interface {
	Init(req *http.Request, resp http.ResponseWriter, pathParmas []string)
	Prepare()
	Finish()

	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
}

func (c *Controller) Init(req *http.Request, resp http.ResponseWriter, pathParams []string) {
	c.Request = req
	c.Response = resp
	c.PathParams = pathParams
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
