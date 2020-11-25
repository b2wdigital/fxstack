package eventbus

import (
	"context"

	"github.com/b2wdigital/fxstack/domain/repository"
)

// NewEvent returns a initialized client
func NewEvent(ctx context.Context) repository.Event {
	return NewPublisher()
}
