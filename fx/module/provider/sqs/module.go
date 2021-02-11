package sqs

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/b2wdigital/fxstack/fx/module/transport/aws"
	provider "github.com/b2wdigital/fxstack/provider/sqs"
	transport "github.com/b2wdigital/fxstack/transport/client/sqs"
	"go.uber.org/fx"
)

var once sync.Once

func EventModule() fx.Option {

	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			aws.AWSModule(),
			fx.Provide(
				sqs.NewFromConfig,
				transport.DefaultOptions,
				transport.NewClient,
				provider.NewEvent,
			),
		)
	})

	return options

}
