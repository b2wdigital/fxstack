package config

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	root             = "aws.default"
	AwsRegion        = root + ".region"
	AwsAccountNumber = root + ".accountNumber"
)

func init() {
	log.Println("getting configurations for fxstack aws")

	giconfig.Add(AwsRegion, "us-east-1", "define aws region")
	giconfig.Add(AwsAccountNumber, "00000000", "define aws account number")
}
