package nats

import (
	"github.com/b2wdigital/fxstack/provider/nats"
	"go.uber.org/fx"
)

var EventModule = fx.Provide(
	nats.NewEvent,
)
