package lambda

import (
	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/fxstack/serverless/lambda"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Handler     cloudevents.Handler
	Middlewares []cloudevents.Middleware `group:"helper"`
	Options     *lambda.Options
}
