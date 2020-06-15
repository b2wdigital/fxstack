package lambda

import (
	"github.com/b2wdigital/fxstack/serverless/lambda"
)

func NewHelper(p Params) (*lambda.Helper, error) {
	return lambda.NewLambdaHelper(p.Handler, p.Middlewares, p.Options)
}
