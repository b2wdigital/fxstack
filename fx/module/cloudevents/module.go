package cloudevents

import (
	"github.com/b2wdigital/fxstack/cloudevents"
	cloudeventsfx "github.com/b2wdigital/fxstack/fx/cloudevents"
	"github.com/b2wdigital/fxstack/fx/module/cloudevents/middleware/eventpublisher"
	"github.com/b2wdigital/fxstack/fx/module/cloudevents/middleware/log"
	"github.com/b2wdigital/fxstack/fx/module/cloudevents/middleware/newrelic"
	gilog "github.com/b2wdigital/goignite/log"
	"go.uber.org/fx"
)

func MiddlewaresModule() fx.Option {

	gilog.Debug("loading cloudevents middleware module")

	return fx.Options(
		log.MiddlewareLogModule(),
		newrelic.MiddlewareNewRelicModule(),
		eventpublisher.MiddlewareLogModule(),
	)
}

func HandlerWrapperModule() fx.Option {

	gilog.Debug("loading cloudevents handler wrapper module")

	return fx.Options(
		fx.Provide(
			cloudevents.DefaultOptions,
			cloudeventsfx.NewHandlerWrapper,
		),
	)
}
