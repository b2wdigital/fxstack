package cloudevents

import (
	"github.com/b2wdigital/fxstack/cloudevents"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Handler     cloudevents.Handler
	Middlewares []cloudevents.Middleware `group:"cloudevents_middlewares"`
	Options     *cloudevents.Options
}
