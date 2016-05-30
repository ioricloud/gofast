// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
	"net/http"
)

type Request struct {
	httpRequest *http.Request
	route       *Route
	parameters  []Parameter
}

type Parameter struct {
	name  string
	value interface{}
}

// Creates a new Request component instance
func NewRequest(req *http.Request, route Route) Request {
	req.ParseForm()

	return Request{req, &route, make([]Parameter, 0)}
}

// Returs HTTP request
func (r *Request) GetHttpRequest() *http.Request {
	return r.httpRequest
}

// Returns current route
func (r *Request) GetRoute() *Route {
	return r.route
}

// Adds a request parameter
func (r *Request) AddParameter(name string, value interface{}) {
	r.parameters = append(r.parameters, Parameter{name, value})
}

// Returns a request parameter from given name
func (r *Request) GetParameter(name string) interface{} {
	var result interface{}

	for _, parameter := range r.parameters {
		if parameter.name == name {
			result = parameter.value
		}
	}

	return result
}

// Returns a POST form value from given name
func (r *Request) GetFormValue(name string) interface{} {
	return r.httpRequest.FormValue(name)
}
