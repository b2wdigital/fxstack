package aws

import (
	giaws "github.com/b2wdigital/goignite/aws/v2"
	"go.uber.org/fx"
)

var awsModuleLoaded = false

func AWSModule() fx.Option {

	if !awsModuleLoaded {
		awsModuleLoaded = true
		return fx.Options(
			fx.Provide(
				giaws.NewDefaultConfig,
			),
		)
	}
	return fx.Options()
}
