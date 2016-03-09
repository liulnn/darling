package darling

import (
	"net/http"
)

type Controller struct {
	Request    *Request
	Response   *Response
	PathParams []string
}

type ControllerInterface interface {
	Init(req *http.Request, resp *Response, pathParmas []string)
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

func (c *Controller) Init(req *Request, resp *Response, pathParams []string) {
	c.Request = req
	c.Response = resp
	c.PathParams = pathParams
}

func (c *Controller) Prepare() {

}

func (c *Controller) Finish() {

}

func (c *Controller) Get() {
	c.Response.StatusCode = 405
	return
}

func (c *Controller) Post() {
	c.Response.StatusCode = 405
	return
}

func (c *Controller) Delete() {
	c.Response.StatusCode = 405
	return
}

func (c *Controller) Put() {
	c.Response.StatusCode = 405
	return
}

func (c *Controller) Head() {
	c.Response.StatusCode = 405
	return
}

func (c *Controller) Patch() {
	c.Response.StatusCode = 405
	return
}

func (c *Controller) Options() {
	c.Response.StatusCode = 405
	return
}
