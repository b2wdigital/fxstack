package nats

import (
	"sync"

	"github.com/b2wdigital/fxstack/provider/nats"
	"go.uber.org/fx"
)

var once sync.Once

func EventModule() fx.Option {

	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			fx.Provide(
				nats.NewEvent,
			),
		)
	})

	return options

}
