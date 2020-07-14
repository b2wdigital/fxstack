package nats

import (
	"context"
	"encoding/json"

	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
	ginats "github.com/b2wdigital/goignite/nats/v1"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
)

type Client struct {
	publisher *ginats.Publisher
}

func NewClient(publisher *ginats.Publisher) *Client {
	return &Client{publisher: publisher}
}

func (p *Client) Publish(ctx context.Context, outs []*v2.Event) (err error) {

	logger := gilog.FromContext(ctx).WithTypeOf(*p)

	logger.Info("publishing to nats")

	for _, out := range outs {

		logger = logger.
			WithField("subject", out.Subject()).
			WithField("id", out.ID())

		var rawMessage []byte

		exts := out.Extensions()

		source, ok := exts["target"]

		if ok {

			s := source.(string)

			if s == "data" {
				var data interface{}

				err = out.DataAs(&data)
				if err != nil {
					return errors.Wrap(err, errors.Internalf("error on data as. %s", err.Error()))
				}

				rawMessage, err = json.Marshal(data)

			} else {
				rawMessage, err = cloudevents.JSONBytes(*out)
			}

		} else {
			rawMessage, err = cloudevents.JSONBytes(*out)
		}

		if err != nil {
			err = errors.Wrap(err, errors.Internalf("error when transforming json into bytes"))
			logger.Error(errors.ErrorStack(err))
			continue
		}

		logger.Info(string(rawMessage))

		msg := &nats.Msg{
			Subject: out.Subject(),
			Data:    rawMessage,
		}

		err = p.publisher.Publish(ctx, msg)
		if err != nil {
			err = errors.Wrap(err, errors.Internalf("unable to publish to nats"))
			logger.Error(errors.ErrorStack(err))
		}

	}

	return nil
}
