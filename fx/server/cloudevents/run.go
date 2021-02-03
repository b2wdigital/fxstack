package cloudevents

import (
	"context"

	"github.com/b2wdigital/fxstack/server/cloudevents"
	gilog "github.com/b2wdigital/goignite/log"
	"go.uber.org/fx"
)

func Run(
	lifecycle fx.Lifecycle,
	helper *cloudevents.Helper,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				helper.Serve()
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
