package main

import (
	"context"

	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/fxstack/serverless/cmd"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
	v2 "github.com/cloudevents/sdk-go/v2"
	logger "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func main() {

	err := cmd.New(
		"sample-serverless",
		"example for serverless application",
		HandlerModule,
		nil,
		nil,
	).
		Run()

	if err != nil {
		gilog.Fatal(err)
	}

}

func HandlerModule() fx.Option {
	return fx.Options(
		fx.Provide(
			NewHandler,
		),
	)
}

type Handler struct {
}

func NewHandler() cloudevents.Handler {
	return &Handler{}
}

func (*Handler) Handle(parentCtx context.Context, in v2.Event) (out *v2.Event, err error) {
	e := v2.NewEvent()
	e.SetSubject("changeme")
	e.SetSource("changeme")
	e.SetType("changeme")
	e.SetExtension("group", "changeme")
	err = e.SetData("", nil)
	if err != nil {
		return nil, errors.Wrap(err, errors.Internalf("unable set out event data"))
	}

	logger.Debug("persisted event")

	return &e, nil
}
