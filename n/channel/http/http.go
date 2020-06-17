package http

import (
	"github.com/b2wdigital/fxstack/n/channel"
	"github.com/b2wdigital/fxstack/n/driver"
)

type HTTP struct {
	drv   driver.HTTP
	route []*Route
}

func New(drv driver.HTTP, route ...*Route) channel.Channel {
	return &HTTP{drv, route}
}

func (c *HTTP) Start() error {
	return nil
}
