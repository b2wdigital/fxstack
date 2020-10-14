package kinesis

import (
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	provk "github.com/b2wdigital/fxstack/provider/kinesis"
	transk "github.com/b2wdigital/fxstack/transport/client/kinesis"
	giaws "github.com/b2wdigital/goignite/aws/v2"
	"go.uber.org/fx"
)

var EventModule = fx.Provide(
	giaws.NewDefaultConfig,
	kinesis.NewFromConfig,
	transk.DefaultOptions,
	transk.NewClient,
	provk.DefaultOptions,
	provk.NewEvent,
)
