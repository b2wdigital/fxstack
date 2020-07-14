package lambda

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/b2wdigital/fxstack/cloudevents"
)

type Helper struct {
	handler *Handler
	//smiddlewares []cloudevents.Middleware
}

func NewLambdaHelper(handler *cloudevents.HandlerWrapper, options *Options) (*Helper, error) {

	h := NewHandler(handler, options)

	return &Helper{
		handler: h,
	}, nil

}

func (h *Helper) Start() {
	lambda.Start(h.handler.Handle)
}
