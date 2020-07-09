package nats

import (
	"context"

	ginats "github.com/b2wdigital/goignite/nats/v1"
	"github.com/b2wdigital/fxstack/cloudevents"
	n "github.com/b2wdigital/fxstack/listener/nats"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Context        context.Context
	Options        *n.Options
	HandlerWrapper *cloudevents.HandlerWrapper
	Queue          *ginats.Queue
}
