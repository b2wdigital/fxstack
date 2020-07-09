package middleware

import (
	"context"
	"time"

	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/fxstack/domain/repository"
	"github.com/b2wdigital/fxstack/wrapper/provider"
	gilog "github.com/b2wdigital/goignite/log"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
)

type EventPublisher struct {
	events repository.Event
}

func NewEventPublisher(events *provider.EventWrapperProvider) cloudevents.Middleware {
	return &EventPublisher{events: events}
}

func (m *EventPublisher) BeforeAll(ctx context.Context, inout []*cloudevents.InOut) (context.Context, error) {
	return ctx, nil
}

func (m *EventPublisher) Before(ctx context.Context, in *v2.Event) (context.Context, error) {
	return ctx, nil
}

func (m *EventPublisher) After(ctx context.Context, in v2.Event, out *v2.Event, err error) (context.Context, error) {
	return ctx, nil
}

func (m *EventPublisher) AfterAll(ctx context.Context, inouts []*cloudevents.InOut) (context.Context, error) {

	logger := gilog.FromContext(ctx).WithTypeOf(*m)

	var outs []*v2.Event

	for _, inout := range inouts {
		if inout.Err != nil {
			logger.Warn("the messages could not be published because one or more messages contain errors.")
			return ctx, nil
		}
	}

	for _, inout := range inouts {

		out := inout.Out
		in := inout.In

		if out != nil {

			id := uuid.New()

			out.SetExtension("parentId", in.ID())
			out.SetID(id.String())
			out.SetTime(time.Now())

			for key, value := range in.Extensions() {
				out.SetExtension(key, value)
			}

			outs = append(outs, out)
		}

	}

	if er := m.events.Publish(ctx, outs); er != nil {
		return ctx, er
	}

	logger.Info("published events")

	return ctx, nil
}

func (m *EventPublisher) Close(ctx context.Context) error {
	return nil
}
