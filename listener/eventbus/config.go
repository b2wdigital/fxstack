package eventbus

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Subjects = "fxstack.listener.eventbus.subjects"
)

func init() {

	log.Println("getting configurations for fxstack listener eventbus")

	giconfig.Add(Subjects, []string{"changeme"}, "eventbus listener subjects")
}

func SubjectsValue() []string {
	return giconfig.Strings(Subjects)
}
