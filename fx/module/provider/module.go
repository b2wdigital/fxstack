package provider

import (
	"github.com/b2wdigital/fxstack/domain/repository"
	"github.com/b2wdigital/fxstack/fx/module/provider/kinesis"
	"github.com/b2wdigital/fxstack/fx/module/provider/mock"
	"github.com/b2wdigital/fxstack/fx/module/provider/nats"
	"github.com/b2wdigital/fxstack/wrapper/provider"
	gilog "github.com/b2wdigital/goignite/log"
	"go.uber.org/fx"
)

func EventModule() fx.Option {

	value := repository.EventProviderValue()

	gilog.Tracef("loading %s event provider module", value)

	var mod fx.Option

	if value == "kinesis" {
		mod = kinesis.EventModule
	} else if value == "nats" {
		mod = nats.EventModule
	} else {
		mod = mock.EventModule
	}

	return fx.Options(mod, fx.Provide(provider.NewEventWrapperProvider))
}
