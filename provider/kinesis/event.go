package kinesis

import (
	"github.com/b2wdigital/fxstack/domain/repository"
	"github.com/b2wdigital/fxstack/transport/client/kinesis"
)

// NewEvent returns a initialized client
func NewEvent(c kinesis.Client, options *Options) repository.Event {
	return NewClient(c, options)
}
