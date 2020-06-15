package repository

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	EventProvider = "fxstack.repository.event.provider"
)

func init() {

	log.Println("getting configurations for fxstack repositories")

	giconfig.Add(EventProvider, "mock", "event provider")
}

func EventProviderValue() string {
	return giconfig.String(EventProvider)
}
