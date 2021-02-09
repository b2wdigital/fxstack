package sqs

import (
	"log"
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Timeout = "fxstack.transport.client.sqs.timeout"
)

func init() {

	log.Println("getting configurations for fxstack sqs transport client")

	giconfig.Add(Timeout, 1*time.Second, "define timeout for sqs client")
}
