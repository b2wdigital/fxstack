package lambda

import (
	"context"

	"github.com/b2wdigital/fxstack/serverless/lambda"
	gilog "github.com/b2wdigital/goignite/log"
	"go.uber.org/fx"
)

func Run(
	lifecycle fx.Lifecycle,
	helper *lambda.Helper,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				helper.Start()
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
