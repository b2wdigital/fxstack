package echo

import (
	"context"

	"github.com/b2wdigital/fxstack/server/rest/echo"
	"go.uber.org/fx"
)

func Start(f func() fx.Option) error {
	return fx.New(
		f(),
		fx.Provide(
			context.Background,
			echo.NewHelper,
		),
		fx.Invoke(Run),
	).Start(context.Background())
}
