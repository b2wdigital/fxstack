package main

import (
	"context"
	"net/http"

	"github.com/b2wdigital/fxstack/server/rest/cmd"
	fxsecho "github.com/b2wdigital/fxstack/server/rest/echo"
	giecho "github.com/b2wdigital/goignite/echo/v4"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {

	err := cmd.New(
		"sample-rest-api",
		"example for rest api",
		HandlerModule,
		ApplicationModule,
		gilogrus.NewLogger,
	).
		Run()

	if err != nil {
		gilog.Fatal(err)
	}

}

func ApplicationModule() fx.Option {
	return fx.Options(
		fx.Provide(
			context.Background,
		),
		fx.Invoke(
			ginewrelic.NewApplication,
		),
	)
}

func HandlerModule() fx.Option {
	return fx.Options(
		fx.Provide(
			Routes,
			NewHandler,
		),
	)
}

func Routes(handler *Handler) []*fxsecho.Route {
	return []*fxsecho.Route{
		fxsecho.NewRouter(http.MethodGet, handler, "/sample"),
	}
}

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (*Handler) Handle(c echo.Context) (err error) {
	return giecho.JSON(c, http.StatusOK, nil, err)
}
