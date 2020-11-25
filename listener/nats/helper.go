package nats

import (
	"context"

	"github.com/b2wdigital/fxstack/cloudevents"
	gilog "github.com/b2wdigital/goignite/log"
	ginats "github.com/b2wdigital/goignite/nats/v1"
)

type Helper struct {
	handler *cloudevents.HandlerWrapper
	queue   string
	q       *ginats.Queue
}

func NewHelper(ctx context.Context, q *ginats.Queue, options *Options,
	handler *cloudevents.HandlerWrapper) (*Helper, error) {

	return &Helper{
		handler: handler,
		queue:   options.Queue,
		q:       q,
	}, nil
}

func (h *Helper) SubscribeAll(subjects []string) {

	for i := range subjects {
		go h.subscribe(context.Background(), subjects[i])
	}

	c := make(chan struct{})
	<-c

}

func (h *Helper) Subscribe(ctx context.Context, subject string) {
	h.subscribe(ctx, subject)

	c := make(chan struct{})
	<-c
}

func (h *Helper) subscribe(ctx context.Context, subject string) {

	logger := gilog.FromContext(ctx).WithTypeOf(*h)

	subscriber := NewSubscriberListener(h.q, h.handler, subject, h.queue)
	subscribe, err := subscriber.Subscribe(ctx)
	if err != nil {
		logger.Error(err)
	}

	if subscribe.IsValid() {
		logger.Infof("nats: subscribed on %s with queue %s", subject, h.queue)
	} else {
		logger.Errorf("nats: not subscribed on %s with queue %s", subject, h.queue)
	}

}
