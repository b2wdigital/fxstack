package nats

import (
	"context"

	"github.com/b2wdigital/fxstack/listener/nats"
	"go.uber.org/fx"
)

func Run(
	lifecycle fx.Lifecycle,
	helper *nats.Helper,
	options *nats.Options,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				helper.SubscribeAll(options.Subjects)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)
}
