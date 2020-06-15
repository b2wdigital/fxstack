package nats

import (
	"context"

	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
	ginats "github.com/b2wdigital/goignite/nats/v1"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
)

// client holds evertything needed to publish a product
type Client struct {
	publisher *ginats.Publisher
}

func NewClient(publisher *ginats.Publisher) *Client {
	return &Client{publisher: publisher}
}

// Publish publishes an array of products on
func (p *Client) Publish(ctx context.Context, outs []*v2.Event) error {

	logger := gilog.FromContext(ctx).WithTypeOf(*p)

	for _, out := range outs {

		rawMessage, err := cloudevents.JSONBytes(*out)
		if err != nil {
			err = errors.Wrap(err, errors.Internalf("error when transforming json into bytes"))
			logger.WithTypeOf(*p).Error(errors.ErrorStack(err))
			continue
		}

		logger.Debug("publishing to nats")
		logger.Debug(string(rawMessage))

		msg := &nats.Msg{
			Subject: out.Subject(),
			Data:    rawMessage,
		}

		err = p.publisher.Publish(ctx, msg)
		if err != nil {
			err = errors.Wrap(err, errors.Internalf("unable to publish to nats"))
			logger.WithTypeOf(*p).Error(errors.ErrorStack(err))
		}

	}

	return nil
}
