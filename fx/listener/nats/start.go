package nats

import (
	"context"

	"go.uber.org/fx"
)

func Start(f func() fx.Option) error {
	return fx.New(
		fx.Provide(context.Background),
		f(),
		HelperModule(),
		fx.Invoke(Run),
	).Start(context.Background())
}
