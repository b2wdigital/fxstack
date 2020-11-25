package nats

import (
	"context"

	"github.com/b2wdigital/fxstack/fx/module/listener/nats"
	"go.uber.org/fx"
)

func Start(f func() fx.Option) error {
	return fx.New(
		fx.Provide(context.Background),
		f(),
		nats.HelperModule(),
		fx.Provide(NewHelper),
		fx.Invoke(Run),
	).Start(context.Background())
}
