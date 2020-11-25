package sns

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
)

// Client knows how to publish on sns
type Client interface {
	Publish(ctx context.Context, input *sns.PublishInput) error
}

// Client holds client and resource name
type client struct {
	client  *sns.Client
	options *Options
}

// NewClient returns a initialized client
func NewClient(c *sns.Client, o *Options) Client {
	return &client{c, o}
}

// Publish publish message on sns
func (c *client) Publish(ctx context.Context, input *sns.PublishInput) error {

	logger := gilog.FromContext(ctx).
		WithTypeOf(*c).
		WithField("subject", input.Subject)

	reqCtx, cancel := context.WithTimeout(context.Background(), c.options.Timeout)
	defer cancel()

	d2 := int64(c.options.Timeout / time.Millisecond)
	logger.WithField("timeout", strconv.FormatInt(d2, 10)).
		Tracef("sending message to sns with timeout: %s", strconv.FormatInt(d2, 10))

	response, err := c.client.Publish(reqCtx, input)
	if err != nil {
		return errors.Wrap(err, errors.New("error publishing message on sns"))
	}

	logger.
		WithField("message_id", *response.MessageId).
		Info("message sent to sns")

	return nil
}
