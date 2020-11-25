package sns

import (
	"log"
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Timeout = "fxstack.transport.client.sns.timeout"
)

func init() {

	log.Println("getting configurations for fxstack sns transport client")

	giconfig.Add(Timeout, 1*time.Second, "define timeout for sns client")
}
