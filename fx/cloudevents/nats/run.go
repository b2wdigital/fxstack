package nats

import (
	"context"

	fsnats "github.com/b2wdigital/fxstack/cloudevents/nats"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
)

func Run(
	lifecycle fx.Lifecycle,
	helper *fsnats.Helper,
	conn *nats.Conn,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				helper.Start(ctx, conn)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger := gilog.FromContext(ctx)
				logger.Info("stopping....")
				return nil
			},
		},
	)
}
