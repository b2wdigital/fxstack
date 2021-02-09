package sqs

import (
	"github.com/b2wdigital/fxstack/domain/repository"
	"github.com/b2wdigital/fxstack/transport/client/sqs"
)

// NewEvent returns a initialized client
func NewEvent(c sqs.Client) repository.Event {
	return NewClient(c)
}
