package mock

import (
	"github.com/b2wdigital/fxstack/provider/mock"
	"go.uber.org/fx"
)

var EventModule = fx.Provide(
	mock.NewEvent,
)
