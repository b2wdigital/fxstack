package cloudevents

import (
	"context"
	"reflect"

	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
	v2 "github.com/cloudevents/sdk-go/v2"
	"golang.org/x/sync/errgroup"
)

type HandlerWrapper struct {
	handler     Handler
	middlewares []Middleware
}

func NewHandlerWrapper(handler Handler, middlewares ...Middleware) *HandlerWrapper {

	if middlewares == nil {
		middlewares = []Middleware{}
	}

	return &HandlerWrapper{handler: handler, middlewares: middlewares}
}

func (h *HandlerWrapper) closeAll(parentCtx context.Context) error {

	logger := gilog.FromContext(parentCtx).WithTypeOf(*h)

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	for _, middleware := range h.middlewares {

		logger.Tracef("executing event middleware %s.AfterAll()", reflect.TypeOf(middleware).String())

		var err error
		err = middleware.Close(ctx)
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

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	for _, middleware := range h.middlewares {

		logger.Tracef("executing event middleware %s.AfterAll()", reflect.TypeOf(middleware).String())

		var err error
		ctx, err = middleware.AfterAll(ctx, inouts)
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

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	for _, middleware := range h.middlewares {

		logger.Tracef("executing event middleware %s.BeforeAll()", reflect.TypeOf(middleware).String())

		var err error
		ctx, err = middleware.BeforeAll(ctx, inouts)
		if err != nil {
			err = errors.Wrap(err,
				errors.Internalf("an error happened when calling BeforeAll() method in %s middleware",
					reflect.TypeOf(middleware).String()))
			return ctx, err
		}
	}

	return ctx, nil
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

	g, gctx := errgroup.WithContext(parentCtx)

	for _, inout := range inouts {

		inout := inout

		g.Go(func() error {

			ctx, cancel := context.WithCancel(parentCtx)
			defer cancel()

			logger := gilog.FromContext(ctx)

			in := inout.In

			l := logger.
				WithField("event.id", in.ID()).
				WithField("event.parentId", in.Extensions()["parentId"]).
				WithField("event.source", in.Source()).
				WithField("event.type", in.Type())

			ctxx := l.ToContext(parentCtx)

			ctxx, err := h.before(ctxx, h.middlewares, in)
			if err != nil {
				return err
			}

			out, err := h.handler.Handle(ctxx, *in)
			if err != nil {
				inout.Err = errors.Wrap(err, errors.Internalf("unable process event"))
			}
			inout.Out = out
			inout.Context = ctxx

			_, err = h.after(ctxx, h.middlewares, *in, out, inout.Err)
			if err != nil {
				return err
			}

			return nil

		})
	}

	if err := g.Wait(); err != nil {
		gilog.FromContext(parentCtx).WithTypeOf(*h).Error(errors.ErrorStack(err))
	}

	gctx.Done()
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
