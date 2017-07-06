package handlers

import (
	"net/http"
)

//Mock pull Mock values for mock server
var Mock bool

//Route ..
type route struct {
	name        string
	pattern     string
	handlerFunc func(http.ResponseWriter, *http.Request)
	handlers    []func(http.Handler) http.Handler
}

var routes = []route{
	route{
		"index",
		"/",
		IndexHandler,
		[]func(http.Handler) http.Handler{
			RecoverHandler,
		},
	},
	route{
		"index",
		"/getmead",
		GetMeAdHandler,
		[]func(http.Handler) http.Handler{
			RecoverHandler,
		},
	},
}

//addRoutes ...
func addRoutes(routes ...route) {
	for _, r := range routes {
		var handler http.Handler
		handler = http.HandlerFunc(r.handlerFunc)
		for i := range r.handlers {
			handler = r.handlers[i](handler)
		}
		http.Handle(r.pattern, handler)
	}
}

//InitHandlers ...
func InitHandlers(mock bool) {
	Mock = mock
	addRoutes(routes...)
}
