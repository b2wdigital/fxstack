package echo

import (
	"context"

	"github.com/b2wdigital/fxstack/server/rest/echo"
	gilog "github.com/b2wdigital/goignite/log"
	"go.uber.org/fx"
)

func Run(
	lifecycle fx.Lifecycle,
	helper *echo.Helper,
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
