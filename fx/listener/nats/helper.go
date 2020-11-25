package nats

import "github.com/b2wdigital/fxstack/listener/nats"

func NewHelper(p Params) (*nats.Helper, error) {
	return nats.NewHelper(p.Context, p.Queue, p.Options, p.HandlerWrapper)
}
