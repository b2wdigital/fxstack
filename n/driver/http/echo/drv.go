package echo

import (
	"github.com/b2wdigital/fxstack/n/channel/http"
	"github.com/b2wdigital/fxstack/n/driver"
)

type drv struct {
}

func (e *drv) Start(route ...*http.Route) error {
	return nil
}

func NewDriver() driver.HTTP {
	return &drv{}
}