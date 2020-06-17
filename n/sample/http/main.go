package main

import (
	"net/http"

	fxshttp "github.com/b2wdigital/fxstack/n/channel/http"
	"github.com/b2wdigital/fxstack/n/cmd"
	fxsecho "github.com/b2wdigital/fxstack/n/driver/http/echo"
	giecho "github.com/b2wdigital/goignite/echo/v4"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/labstack/echo/v4"
)

func main() {

	echodrv := fxsecho.NewDriver()
	httpChan := fxshttp.New(echodrv, fxshttp.NewRouter("GET", NewHelloHandler(), "/hello"))

	co := cmd.New()
	co.WithChan(httpChan)
}

type HelloHandler struct {
}

func NewHelloHandler() fxsecho.Handler {
	return &HelloHandler{}
}

func (*HelloHandler) Handle(c echo.Context) (err error) {

	logger := gilog.FromContext(c.Request().Context())

	logger.Info("example of request")

	return giecho.JSON(c, http.StatusOK, nil, err)
}
