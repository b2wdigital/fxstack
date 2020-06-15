package middleware

import (
	"context"
	"encoding/json"

	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
	v2 "github.com/cloudevents/sdk-go/v2"
)

type Log struct {
}

func NewLog() cloudevents.Middleware {
	return &Log{}
}

func (m *Log) Close(ctx context.Context) error {
	return nil
}

func (m *Log) BeforeAll(ctx context.Context, inout []*cloudevents.InOut) (context.Context, error) {
	return ctx, nil
}

func (m *Log) Before(ctx context.Context, in *v2.Event) (context.Context, error) {
	logger := gilog.FromContext(ctx).WithTypeOf(*m)

	logger.Info("received event")

	j, _ := json.Marshal(in)
	logger.Info(string(j))

	return ctx, nil
}

func (m *Log) After(ctx context.Context, in v2.Event, out *v2.Event, err error) (context.Context, error) {
	logger := gilog.FromContext(ctx).WithTypeOf(*m)

	if out != nil && err == nil {

		logger.Info("returning event")

		j, _ := json.Marshal(out)
		logger.Info(string(j))

	}

	if err != nil {
		logger.Error(errors.ErrorStack(err))
	}

	return ctx, nil
}

func (m *Log) AfterAll(ctx context.Context, inout []*cloudevents.InOut) (context.Context, error) {
	return ctx, nil
}
