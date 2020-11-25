package eventbus

import (
	"context"

	"github.com/b2wdigital/fxstack/cloudevents"
	gilog "github.com/b2wdigital/goignite/log"
)

type Helper struct {
	handler *cloudevents.HandlerWrapper
}

func NewHelper(ctx context.Context, handler *cloudevents.HandlerWrapper) (*Helper, error) {
	return &Helper{handler: handler}, nil
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

	subscriber := NewSubscriberListener(h.handler, subject)
	err := subscriber.Subscribe(ctx)
	if err != nil {
		logger.Error(err)
	}
}
