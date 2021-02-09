package provider

import (
	"sync"

	"github.com/b2wdigital/fxstack/domain/repository"
	"github.com/b2wdigital/fxstack/fx/module/provider/eventbus"
	"github.com/b2wdigital/fxstack/fx/module/provider/kinesis"
	"github.com/b2wdigital/fxstack/fx/module/provider/mock"
	"github.com/b2wdigital/fxstack/fx/module/provider/nats"
	"github.com/b2wdigital/fxstack/fx/module/provider/sns"
	"github.com/b2wdigital/fxstack/fx/module/provider/sqs"
	"github.com/b2wdigital/fxstack/wrapper/provider"
	gilog "github.com/b2wdigital/goignite/log"
	"go.uber.org/fx"
)

var once sync.Once

func EventModule() fx.Option {

	options := fx.Options()

	once.Do(func() {

		value := repository.EventProviderValue()

		gilog.Tracef("loading %s event provider module", value)

		var mod fx.Option

		if value == "kinesis" {
			mod = kinesis.EventModule()
		} else if value == "nats" {
			mod = nats.EventModule()
		} else if value == "eventbus" {
			mod = eventbus.EventModule()
		} else if value == "sns" {
			mod = sns.EventModule()
		} else if value == "sqs" {
			mod = sqs.EventModule()
		} else {
			mod = mock.EventModule()
		}

		options = fx.Options(mod, fx.Provide(provider.NewEventWrapperProvider))

	})

	return options
}
