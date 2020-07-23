package nats

import (
	"context"
	"encoding/json"

	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/goignite/errors"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	ginats "github.com/b2wdigital/goignite/nats/v1"
	v2 "github.com/cloudevents/sdk-go/v2"
)

type SubscriberListener struct {
	handler *cloudevents.HandlerWrapper
	subject string
}

func NewSubscriberListener(q *ginats.Queue, handler *cloudevents.HandlerWrapper, subject string,
	queue string) *SubscriberListener {
	return &SubscriberListener{
		handler: handler,
		subject: subject,
	}
}

func (l *SubscriberListener) Subscribe(ctx context.Context) error {
	return gieventbus.Subscribe(l.subject, l.h)
}

func (l *SubscriberListener) SubscribeOnce(ctx context.Context) error {
	return gieventbus.SubscribeOnce(l.subject, l.h)
}

func (l *SubscriberListener) h(event []byte) {

	in := v2.NewEvent()
	err := json.Unmarshal(event, &in)
	if err != nil {

		var data interface{}

		if err := json.Unmarshal(event, &data); err != nil {
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
		WithField("subject", l.subject)

	ctx := logger.ToContext(context.Background())

	var inouts []*cloudevents.InOut

	inouts = append(inouts, &cloudevents.InOut{In: &in})

	err = l.handler.Process(ctx, inouts)
	if err != nil {
		logger.Error(errors.ErrorStack(err))
	}

}
