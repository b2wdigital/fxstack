package lambda

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/b2wdigital/fxstack/cloudevents"
)

type Helper struct {
	handler *Handler
	//smiddlewares []cloudevents.Middleware
}

func NewLambdaHelper(handler cloudevents.Handler, middlewares []cloudevents.Middleware,
	options *Options) (*Helper, error) {

	h := NewHandler(handler, middlewares, options)

	return &Helper{
		handler: h,
	}, nil

}

func (h *Helper) Start() {
	lambda.Start(h.handler.Handle)
}
