package repository

import (
	"context"

	"github.com/cloudevents/sdk-go/v2" //nolint
)

// Event knows how to publish products
type Event interface {
	Publish(context.Context, []*v2.Event) error
}
