package nats

import (
	"context"

	"github.com/b2wdigital/fxstack/cloudevents"
	n "github.com/b2wdigital/fxstack/listener/nats"
	ginats "github.com/b2wdigital/goignite/nats/v1"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Context     context.Context
	Options     *n.Options
	Handler     cloudevents.Handler
	Middlewares []cloudevents.Middleware `group:"helper"`
	Queue       *ginats.Queue
}
