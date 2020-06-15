package middleware

import (
	"context"

	"github.com/b2wdigital/fxstack/cloudevents"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	v2 "github.com/cloudevents/sdk-go/v2"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

type NewRelic struct {
}

func NewNewRelic() cloudevents.Middleware {
	return &NewRelic{}
}

func (m *NewRelic) BeforeAll(ctx context.Context, inout []*cloudevents.InOut) (context.Context, error) {

	txn := ginewrelic.Application().StartTransaction(cloudevents.MiddlewareNewRelicTxNameValue())

	c := nr.NewContext(ctx, txn)

	return c, nil
}

func (m *NewRelic) Before(parentCtx context.Context, in *v2.Event) (context.Context, error) {

	txn := nr.FromContext(parentCtx).NewGoroutine()

	ctx := nr.NewContext(parentCtx, txn)

	return ctx, nil
}

func (m *NewRelic) After(parentCtx context.Context, in v2.Event, out *v2.Event, err error) (context.Context, error) {

	txn := nr.FromContext(parentCtx)

	if err != nil {
		if txn != nil {
			txn.NoticeError(err)
		}
	}

	return parentCtx, nil
}

func (m *NewRelic) AfterAll(ctx context.Context, inout []*cloudevents.InOut) (context.Context, error) {
	return ctx, nil
}

func (m *NewRelic) Close(ctx context.Context) error {
	txn := nr.FromContext(ctx)
	defer txn.End()

	return nil
}
