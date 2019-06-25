package darling

import (
	"net/http"
	"regexp"
)

type ControllerInfo struct {
	regex      *regexp.Regexp
	controller ControllerInterface
}

type ControllerRegister struct {
	routers []*ControllerInfo
}

func NewControllerRegister() *ControllerRegister {
	return &ControllerRegister{
		routers: make([]*ControllerInfo, 0),
	}
}

func (p *ControllerRegister) Add(pattern string, c ControllerInterface) {

	regex, regexErr := regexp.Compile(pattern)
	if regexErr != nil {

		panic(regexErr)
		return
	}
	p.routers = append(p.routers, &ControllerInfo{regex: regex, controller: c})
}

func (p *ControllerRegister) Extend(q *ControllerRegister) {
	p.routers = append(p.routers, q.routers...)
}

func (p *ControllerRegister) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var started bool
	requestPath := r.URL.Path
	req := &Request{In: r}
	resp := &Response{Out: rw, Header: make(map[string]string)}

	for _, route := range p.routers {

		if !route.regex.MatchString(requestPath) {
			continue
		}
		matches := route.regex.FindStringSubmatch(requestPath)
		if len(matches[0]) != len(requestPath) {
			continue
		}

		route.controller.Init(req, resp, matches[1:])
		route.controller.Prepare()
		switch r.Method {
		case "OPTIONS":
			route.controller.Options()
		case "HEAD":
			route.controller.Post()
		case "GET":
			route.controller.Get()
		case "POST":
			route.controller.Post()
		case "PUT":
			route.controller.Put()
		case "DELETE":
			route.controller.Delete()
		case "PATCH":
			route.controller.Patch()
		default:
			http.Error(rw, "Method Not Allowed", 405)
		}
		route.controller.Finish()
		resp.Write()
		started = true
		break
	}

	if started == false {
		http.NotFound(rw, r)
	}
}
