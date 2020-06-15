package nats

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Subjects = "fxstack.listener.nats.subjects"
	Queue    = "fxstack.listener.nats.queue"
)

func init() {

	log.Println("getting configurations for fxstack listener nats")

	giconfig.Add(Subjects, []string{"changeme"}, "nats listener subjects")
	giconfig.Add(Queue, "changeme", "nats listener queue")
}

func SubjectsValue() []string {
	return giconfig.Strings(Subjects)
}

func QueueValue() string {
	return giconfig.String(Queue)
}
