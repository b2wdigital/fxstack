package echo

import (
	"context"
	"net/http"

	giecho "github.com/b2wdigital/goignite/echo/v4"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/labstack/echo/v4"
)

type Helper struct {
	echo   *echo.Echo
	ctx    context.Context
	routes []*Route
}

func NewHelper(ctx context.Context, routes []*Route) (*Helper, error) {

	ec := giecho.Start(ctx)

	for _, route := range routes {
		addRoute(ctx, route, ec)
	}

	return &Helper{
		echo:   ec,
		ctx:    ctx,
		routes: routes,
	}, nil
}

func addRoute(ctx context.Context, route *Route, e *echo.Echo) {

	var r func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

	logger := gilog.FromContext(ctx)

	switch method := route.Method(); method {
	case http.MethodPost:
		r = e.POST
	case http.MethodPut:
		r = e.PUT
	case http.MethodDelete:
		r = e.DELETE
	case http.MethodHead:
		r = e.HEAD
	default:
		r = e.GET
	}

	logger.Infof("configuring app router on %s for method %s", route.Route(), route.Method())

	r(route.Route(), route.Handler().Handle, route.Middleware()...)
}

func (h *Helper) Serve() {
	giecho.Serve(h.ctx)
}