package newrelic

import (
	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/fxstack/cloudevents/middleware"
	"go.uber.org/fx"
)

func MiddlewareNewRelicModule() fx.Option {

	if cloudevents.MiddlewareNewRelicEnabledValue() {
		return fx.Provide(
			fx.Annotated{
				Group:  "helper",
				Target: middleware.NewNewRelic,
			},
		)
	}
	return fx.Provide()

}
