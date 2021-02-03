package cloudevents

import (
	"context"

	"github.com/b2wdigital/fxstack/server/cloudevents"
	"go.uber.org/fx"
)

func Start(f func() fx.Option) error {
	return fx.New(
		f(),
		fx.Provide(
			context.Background,
			cloudevents.NewHelper,
		),
		fx.Invoke(Run),
	).Start(context.Background())
}
