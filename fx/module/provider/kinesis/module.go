package kinesis

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/b2wdigital/fxstack/fx/module/transport/aws"
	provk "github.com/b2wdigital/fxstack/provider/kinesis"
	transk "github.com/b2wdigital/fxstack/transport/client/kinesis"
	"go.uber.org/fx"
)

var once sync.Once

func EventModule() fx.Option {

	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			aws.AWSModule(),
			fx.Provide(
				kinesis.NewFromConfig,
				transk.DefaultOptions,
				transk.NewClient,
				provk.DefaultOptions,
				provk.NewEvent,
			),
		)
	})

	return options

}
