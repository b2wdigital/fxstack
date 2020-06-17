package http

import (
	"github.com/b2wdigital/fxstack/n/driver/http/echo"
)

type Route struct {
	method   string
	handler  echo.Handler
	route    string
	httpCode int
}

func (r *Route) Route() string {
	return r.route
}

func (r *Route) Handler() echo.Handler {
	return r.handler
}

func (r *Route) Method() string {
	return r.method
}

func NewRouter(method string, handler echo.Handler, route string) *Route {
	return &Route{method: method, handler: handler, route: route}
}
