package sqs

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
)

// Client knows how to publish on sqs
type Client interface {
	Publish(ctx context.Context, input *sqs.SendMessageInput) error
}

// Client holds client and resource name
type client struct {
	client  *sqs.Client
	options *Options
}

// NewClient returns a initialized client
func NewClient(c *sqs.Client, o *Options) Client {
	return &client{c, o}
}

// Publish publish message on sns
func (c *client) Publish(ctx context.Context, input *sqs.SendMessageInput) error {

	logger := gilog.FromContext(ctx).
		WithTypeOf(*c).
		WithField("subject", input.QueueUrl)

	reqCtx, cancel := context.WithTimeout(context.Background(), c.options.Timeout)
	defer cancel()

	d2 := int64(c.options.Timeout / time.Millisecond)
	logger.WithField("timeout", strconv.FormatInt(d2, 10)).
		Tracef("sending message to sqs with timeout: %s", strconv.FormatInt(d2, 10))

	response, err := c.client.SendMessage(reqCtx, input)
	if err != nil {
		return errors.Wrap(err, errors.New("error sending message to sqs"))
	}

	logger.
		WithField("message_id", *response.MessageId).
		Info("message sent to sqs")

	return nil
}
