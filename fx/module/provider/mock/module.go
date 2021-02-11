package mock

import (
	"sync"

	"github.com/b2wdigital/fxstack/provider/mock"
	"go.uber.org/fx"
)

var once sync.Once

func EventModule() fx.Option {

	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			fx.Provide(
				mock.NewEvent,
			),
		)
	})

	return options

}
