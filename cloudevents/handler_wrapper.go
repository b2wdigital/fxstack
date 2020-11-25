package cloudevents

import (
	"context"
	"reflect"

	"github.com/b2wdigital/fxstack/util"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
	v2 "github.com/cloudevents/sdk-go/v2"
)

type HandlerWrapper struct {
	handler     Handler
	middlewares []Middleware
	options     *Options
}

func NewHandlerWrapper(handler Handler, options *Options, middlewares ...Middleware) *HandlerWrapper {

	if middlewares == nil {
		middlewares = []Middleware{}
	}

	return &HandlerWrapper{handler: handler, middlewares: middlewares, options: options}
}

func (h *HandlerWrapper) closeAll(parentCtx context.Context) error {

	logger := gilog.FromContext(parentCtx).WithTypeOf(*h)

	for _, middleware := range h.middlewares {

		logger.Tracef("executing event middleware %s.Close()", reflect.TypeOf(middleware).String())

		var err error
		err = middleware.Close(parentCtx)
		if err != nil {
			err = errors.Wrap(err, errors.Internalf("an error happened when calling Close() method in %s middleware",
				reflect.TypeOf(middleware).String()))
			return err
		}
	}

	return nil
}

func (h *HandlerWrapper) afterAll(parentCtx context.Context, inouts []*InOut) error {

	logger := gilog.FromContext(parentCtx).WithTypeOf(*h)

	for _, middleware := range h.middlewares {

		logger.Tracef("executing event middleware %s.AfterAll()", reflect.TypeOf(middleware).String())

		var err error
		parentCtx, err = middleware.AfterAll(parentCtx, inouts)
		if err != nil {
			err = errors.Wrap(err,
				errors.Internalf("an error happened when calling AfterAll() method in %s middleware",
					reflect.TypeOf(middleware).String()))
			return err
		}
	}

	return nil
}

func (h *HandlerWrapper) beforeAll(parentCtx context.Context, inouts []*InOut) (context.Context, error) {

	logger := gilog.FromContext(parentCtx).WithTypeOf(*h)

	for _, middleware := range h.middlewares {

		logger.Tracef("executing event middleware %s.BeforeAll()", reflect.TypeOf(middleware).String())

		var err error
		parentCtx, err = middleware.BeforeAll(parentCtx, inouts)
		if err != nil {
			err = errors.Wrap(err,
				errors.Internalf("an error happened when calling BeforeAll() method in %s middleware",
					reflect.TypeOf(middleware).String()))
			return parentCtx, err
		}
	}

	return parentCtx, nil
}

func (h *HandlerWrapper) Process(parentCtx context.Context, inouts []*InOut) (err error) {

	logger := gilog.FromContext(parentCtx).WithTypeOf(*h)

	parentCtx, err = h.beforeAll(parentCtx, inouts)
	if err != nil {
		logger.Error(errors.ErrorStack(err))
		return err
	}

	h.handleAll(parentCtx, inouts)

	err = h.afterAll(parentCtx, inouts)
	if err != nil {
		logger.Error(errors.ErrorStack(err))
		return err
	}

	err = h.closeAll(parentCtx)
	if err != nil {
		logger.Error(errors.ErrorStack(err))
		return err
	}

	for _, inout := range inouts {
		if inout.Err != nil {
			return err
		}
	}

	return nil
}

func (h *HandlerWrapper) handleAll(parentCtx context.Context, inouts []*InOut) {

	for _, inout := range inouts {
		logger := gilog.FromContext(parentCtx)

		in := inout.In
		if in == nil {
			logger.Warn("discarding inout.In == nil")
			continue
		}

		l := logger.
			WithField("event.id", in.ID()).
			WithField("event.parentId", in.Extensions()["parentId"]).
			WithField("event.source", in.Source()).
			WithField("event.type", in.Type())

		l.Info("event received")

		if inout.Err != nil {
			l.WithField("cause", inout.Err.Error()).Warn("discarding message due to error")
			inout.Err = nil
			continue
		}

		hasToDiscardEvent := util.StringSliceContains(h.options.IDsToDiscard, in.ID())
		if hasToDiscardEvent {
			l.Warn("discarding event due to feature flag")
			continue
		}

		ctxx := l.ToContext(parentCtx)

		ctxx, err := h.before(ctxx, h.middlewares, in)
		if err != nil {
			l.WithField("cause", err.Error()).Warn("could not execute h.before")
		}

		out, err := h.handler.Handle(ctxx, *in)
		if err != nil {
			inout.Err = errors.Wrap(err, errors.Internalf("unable process event"))
		}
		inout.Out = out
		inout.Context = ctxx

		_, err = h.after(ctxx, h.middlewares, *in, out, inout.Err)
		if err != nil {
			l.WithField("cause", err.Error()).Warn("could not execute h.after")
		}
	}
}

func (h *HandlerWrapper) before(ctx context.Context, middlewares []Middleware, in *v2.Event) (context.Context, error) {

	logger := gilog.FromContext(ctx).WithTypeOf(*h)

	var err error

	for _, middleware := range middlewares {

		logger.Tracef("executing event middleware %s.Before()", reflect.TypeOf(middleware).String())

		ctx, err = middleware.Before(ctx, in)
		if err != nil {
			return ctx, errors.Wrap(err,
				errors.Internalf("an error happened when calling Before() method in %s middleware",
					reflect.TypeOf(middleware).String()))
		}
	}

	return ctx, nil
}

func (h *HandlerWrapper) after(ctx context.Context, middlewares []Middleware, in v2.Event, out *v2.Event,
	err error) (context.Context, error) {

	logger := gilog.FromContext(ctx).WithTypeOf(*h)

	var er error

	for _, middleware := range middlewares {

		logger.Tracef("executing event middleware %s.After()", reflect.TypeOf(middleware).String())

		ctx, er = middleware.After(ctx, in, out, err)
		if er != nil {
			return ctx, errors.Wrap(err,
				errors.Internalf("an error happened when calling After() method in %s middleware",
					reflect.TypeOf(middleware).String()))
		}
	}

	return ctx, nil
}
