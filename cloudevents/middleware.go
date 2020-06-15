package cloudevents

import (
	"context"

	v2 "github.com/cloudevents/sdk-go/v2"
)

type Middleware interface {
	BeforeAll(ctx context.Context, inout []*InOut) (context.Context, error)
	Before(ctx context.Context, in *v2.Event) (context.Context, error)
	After(ctx context.Context, in v2.Event, out *v2.Event, err error) (context.Context, error)
	AfterAll(ctx context.Context, inout []*InOut) (context.Context, error)
	Close(ctx context.Context) error
}
