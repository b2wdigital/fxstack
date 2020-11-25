package sns

import (
	"github.com/b2wdigital/fxstack/domain/repository"
	"github.com/b2wdigital/fxstack/transport/client/sns"
)

// NewEvent returns a initialized client
func NewEvent(c sns.Client) repository.Event {
	return NewClient(c)
}
