package log

import (
	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/fxstack/cloudevents/middleware"
	"go.uber.org/fx"
)

func MiddlewareLogModule() fx.Option {

	if cloudevents.MiddlewareLogEnabledValue() {

		return fx.Provide(

			fx.Annotated{
				Group:  "helper",
				Target: middleware.NewLog,
			},
		)

	}
	return fx.Provide()

}
