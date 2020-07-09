package cloudevents

import "github.com/b2wdigital/fxstack/cloudevents"

func NewHandlerWrapper(p Params) *cloudevents.HandlerWrapper {
	return cloudevents.NewHandlerWrapper(p.Handler, p.Options, p.Middlewares...)
}
