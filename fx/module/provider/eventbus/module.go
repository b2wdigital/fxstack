package eventbus

import (
	"github.com/b2wdigital/fxstack/provider/eventbus"
	"go.uber.org/fx"
)

var EventModule = fx.Provide(
	eventbus.NewEvent,
)
