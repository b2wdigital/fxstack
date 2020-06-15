package lambda

import (
	"github.com/b2wdigital/fxstack/fx/module/cloudevents"
	"github.com/b2wdigital/fxstack/serverless/lambda"
	gilog "github.com/b2wdigital/goignite/log"
	"go.uber.org/fx"
)

func HelperModule() fx.Option {

	gilog.Debug("loading lambda helper module")

	return fx.Options(
		cloudevents.MiddlewaresModule(),
		fx.Provide(
			lambda.DefaultOptions,
			NewHelper,
		),
	)
}
