package darling

import (
	"net/http"
	"regexp"
)

type ControllerInfo struct {
	regex      *regexp.Regexp
	controller ControllerInterface
}

type ControllerRegistor struct {
	routers []*ControllerInfo
}

func NewControllerRegistor() *ControllerRegistor {
	return &ControllerRegistor{
		routers: make([]*ControllerInfo, 0),
	}
}

func (p *ControllerRegistor) Add(pattern string, c ControllerInterface) {

	regex, regexErr := regexp.Compile(pattern)
	if regexErr != nil {

		panic(regexErr)
		return
	}
	p.routers = append(p.routers, &ControllerInfo{regex: regex, controller: c})
}

func (p *ControllerRegistor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var started bool
	requestPath := r.URL.Path
	resp := &Response{Out: rw, Header: make(map[string]string)}

	for _, route := range p.routers {

		if !route.regex.MatchString(requestPath) {
			continue
		}
		matches := route.regex.FindStringSubmatch(requestPath)
		if len(matches[0]) != len(requestPath) {
			continue
		}

		route.controller.Init(r, resp, matches[1:])
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
