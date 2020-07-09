package rest

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	liveEnabled = "fxstack.server.rest.live.enabled"
	livePort    = "fxstack.server.rest.live.port"
)

func init() {

	log.Println("getting configurations for fxstack server rest")

	giconfig.Add(liveEnabled, false, "live enable/disable")
	giconfig.Add(livePort, 8081, "live port")
}

func LiveEnabled() bool {
	return giconfig.Bool(liveEnabled)
}

func LivePort() int {
	return giconfig.Int(livePort)
}
