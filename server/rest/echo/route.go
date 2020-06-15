package echo

import (
	"github.com/labstack/echo/v4"
)

type Route struct {
	method     string
	handler    Handler
	route      string
	middleware []echo.MiddlewareFunc
}

func (r *Route) Route() string {
	return r.route
}

func (r *Route) Handler() Handler {
	return r.handler
}

func (r *Route) Method() string {
	return r.method
}

func (r *Route) Middleware() []echo.MiddlewareFunc {
	return r.middleware
}

func NewRouter(method string, handler Handler, route string, middleware ...echo.MiddlewareFunc) *Route {
	return &Route{method: method, handler: handler, route: route, middleware: middleware}
}
