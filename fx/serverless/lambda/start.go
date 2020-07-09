package lambda

import (
	"context"

	"github.com/b2wdigital/fxstack/fx/module/lambda"
	"go.uber.org/fx"
)

func Start(f func() fx.Option) error {
	return fx.New(
		fx.Provide(context.Background),
		f(),
		lambda.HelperModule(),
		fx.Invoke(Run),
	).Start(context.Background())
}
