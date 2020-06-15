package nats

import (
	"context"

	"github.com/b2wdigital/fxstack/domain/repository"
	ginats "github.com/b2wdigital/goignite/nats/v1"
)

// NewEvent returns a initialized client
func NewEvent(ctx context.Context) repository.Event {
	publisher, _ := ginats.NewDefaultPublisher(ctx)
	return NewClient(publisher)
}
