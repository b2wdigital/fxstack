package eventpublisher

import (
	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/fxstack/cloudevents/middleware"
	"github.com/b2wdigital/fxstack/fx/module/provider"
	"go.uber.org/fx"
)

func MiddlewareLogModule() fx.Option {

	if cloudevents.MiddlewareEventPublisherEnabledValue() {

		return fx.Options(
			provider.EventModule(),
			fx.Provide(
				fx.Annotated{
					Group:  "cloudevents_middlewares",
					Target: middleware.NewEventPublisher,
				},
			),
		)

	}

	return fx.Provide()

}
