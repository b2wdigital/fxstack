package nats

import (
	"context"

	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/goignite/cloudevents/v2/nats/receiver"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/nats-io/nats.go"
)

type Helper struct {
	handler  cloudevents.Handler
	subjects []string
}

func NewHelper(handler cloudevents.Handler, subjects []string, middlewares []cloudevents.Middleware) (*Helper, error) {

	gilog.Debugf("loading %v middlewares on cloudevents helper", len(middlewares))

	// h := cloudevents.NewHandlerWrapper(handler, middlewares...)

	return &Helper{
		handler:  handler,
		subjects: subjects,
	}, nil

}

func (h *Helper) Start(ctx context.Context, conn *nats.Conn) {

	logger := gilog.FromContext(ctx)

	options, err := receiver.DefaultOptions()
	if err != nil {
		logger.Fatal(err)
	}

	options.Subjects = h.subjects
	receiver.StartConsumer(ctx, conn, h.handler.Handle, options)
}
