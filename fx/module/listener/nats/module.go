package nats

import (
	"github.com/b2wdigital/fxstack/fx/module/cloudevents"
	"github.com/b2wdigital/fxstack/listener/nats"
	gilog "github.com/b2wdigital/goignite/log"
	ginats "github.com/b2wdigital/goignite/nats/v1"
	"go.uber.org/fx"
)

func HelperModule() fx.Option {

	gilog.Debug("loading listener nats module")

	return fx.Options(
		cloudevents.MiddlewaresModule(),
		fx.Provide(
			ginats.NewDefaultQueue,
			nats.DefaultOptions,
			NewHelper,
		),
	)
}
