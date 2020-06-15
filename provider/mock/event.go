package mock

import "github.com/b2wdigital/fxstack/domain/repository"

// NewEvent returns a initialized mock
func NewEvent() repository.Event {
	return NewMock()
}
