package kinesis

import (
	"log"
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	ReconnectWait = "fxstack.transport.client.kinesis.timeout"
)

func init() {

	log.Println("getting configurations for fxstack kinesis transport client")

	giconfig.Add(ReconnectWait, 1*time.Second, "define timeout for kinesis client")
}
