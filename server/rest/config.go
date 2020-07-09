package rest

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	liveEnabled = "fxstack.server.rest.live.enabled"
	livePort    = "fxstack.server.rest.live.port"
	livePath    = "fxstack.server.rest.live.path"
)

func init() {

	log.Println("getting configurations for fxstack server rest")

	giconfig.Add(liveEnabled, false, "live enable/disable")
	giconfig.Add(livePort, 8081, "live port")
	giconfig.Add(livePath, "/live", "live path")
}

func LiveEnabled() bool {
	return giconfig.Bool(liveEnabled)
}

func LivePort() int {
	return giconfig.Int(livePort)
}

func LivePath() string {
	return giconfig.String(livePath)
}
