package lambda

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Skip = "fxstack.serverless.lambda.skip"
)

func init() {

	log.Println("getting configurations for fxstack serverless lambda")

	giconfig.Add(Skip, false, "skip all triggers")
}

func SkipValue() bool {
	return giconfig.Bool(Skip)
}
