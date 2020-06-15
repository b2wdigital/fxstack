package nats

import (
	"context"
	"encoding/json"

	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
	ginats "github.com/b2wdigital/goignite/nats/v1"
	v2 "github.com/cloudevents/sdk-go/v2"
	n "github.com/nats-io/nats.go"
)

type SubscriberListener struct {
	q       *ginats.Queue
	handler *cloudevents.HandlerWrapper
	subject string
	queue   string
}

func NewSubscriberListener(q *ginats.Queue, handler *cloudevents.HandlerWrapper, subject string,
	queue string) *SubscriberListener {
	return &SubscriberListener{
		q:       q,
		handler: handler,
		subject: subject,
		queue:   queue,
	}
}

func (l *SubscriberListener) Subscribe(ctx context.Context) (*n.Subscription, error) {
	return l.q.Subscribe(l.subject, l.queue, l.h)
}

func (l *SubscriberListener) h(msg *n.Msg) {

	in := v2.NewEvent()
	err := json.Unmarshal(msg.Data, &in)
	if err != nil {

		var data interface{}

		if err := json.Unmarshal(msg.Data, &data); err != nil {
			gilog.Errorf("could not decode nats record. %s", err.Error())
		} else {
			err := in.SetData("", data)
			if err != nil {
				gilog.Errorf("could set data from nats record. %s", err.Error())
				return
			}
		}

	}

	logger := gilog.WithTypeOf(*l).
		WithField("subject", l.subject).
		WithField("queue", l.queue)

	ctx := logger.ToContext(context.Background())

	var inouts []*cloudevents.InOut

	inouts = append(inouts, &cloudevents.InOut{In: &in})

	err = l.handler.Process(ctx, inouts)
	if err != nil {
		logger.Error(errors.ErrorStack(err))
	}

}
