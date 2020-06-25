package kinesis

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
)

// Client knows how to bulkpublish on kinesis
type Client interface {
	BulkPublish(ctx context.Context, messages []kinesis.PutRecordsRequestEntry, resource string) error
	Publish(ctx context.Context, input *kinesis.PutRecordInput) error
}

// Client holds client and resource name
type client struct {
	client  *kinesis.Client
	options *Options
}

// NewClient returns a initialized client
func NewClient(c *kinesis.Client, o *Options) Client {
	return &client{c, o}
}

// Publish publish message on kinesis
func (c *client) Publish(ctx context.Context, input *kinesis.PutRecordInput) error {

	logger := gilog.FromContext(ctx).
		WithTypeOf(*c).
		WithField("resource", input.StreamName).
		WithField("partition_key", input.PartitionKey)

	request := c.client.PutRecordRequest(input)

	reqCtx, cancel := context.WithTimeout(context.Background(), c.options.Timeout)
	defer cancel()

	d2 := int64(c.options.Timeout / time.Millisecond)
	logger.WithField("timeout", strconv.FormatInt(d2, 10)).
		Tracef("sending message to kinesis with timeout: %s", strconv.FormatInt(d2, 10))

	response, err := request.Send(reqCtx)
	if err != nil {
		return errors.Wrap(err, errors.New("error publishing message on kinesis"))
	}

	logger.
		WithField("sequence_number", *response.SequenceNumber).
		WithField("shard_id", *response.ShardId).
		Debug("message sent to kinesis")

	return nil
}

// BulkPublish publishes an array of messages on kinesis
func (c *client) BulkPublish(ctx context.Context, messages []kinesis.PutRecordsRequestEntry, resource string) error {

	logger := gilog.FromContext(ctx).
		WithTypeOf(*c).
		WithField("resource", resource)

	input := c.buildPutRecordsInput(messages, resource)

	request := c.client.PutRecordsRequest(input)

	reqCtx, cancel := context.WithTimeout(context.Background(), c.options.Timeout)
	defer cancel()

	d2 := int64(c.options.Timeout / time.Millisecond)
	logger.WithField("timeout", strconv.FormatInt(d2, 10)).
		Debugf("sending bulk message to kinesis with timeout: %s", strconv.FormatInt(d2, 10))

	response, err := request.Send(reqCtx)
	if err != nil {
		return errors.Wrap(err, errors.New("error publishing message on kinesis"))
	}

	if *response.FailedRecordCount > int64(0) {

		logger.Warnf("Error on publishing bulk messages. total errors: %v / %v",
			*response.FailedRecordCount, len(messages))

	}

	var retry []kinesis.PutRecordsRequestEntry

	for i, r := range response.PutRecordsOutput.Records {

		if r.ErrorMessage != nil {
			logger.
				WithField("cause", r.ErrorMessage).
				WithField("code", r.ErrorCode).
				Warn("error in kinesis bulk record")
			retry = append(retry, messages[i])
			continue
		}

		logger.
			WithField("sequence_number", *r.SequenceNumber).
			WithField("shard_id", *r.ShardId).
			Debug("message sent to kinesis")

	}

	if len(retry) > 0 {

		logger.Warnf("Retrying publish %v messages", len(retry))

		err := c.BulkPublish(ctx, retry, resource)
		if err != nil {
			logger.WithField("cause", err.Error()).Warn("error in kinesis bulk record")
			return err
		}

	}

	return nil
}

func (c *client) buildPutRecordsInput(messages []kinesis.PutRecordsRequestEntry,
	resource string) *kinesis.PutRecordsInput {
	return &kinesis.PutRecordsInput{
		Records:    messages,
		StreamName: aws.String(resource),
	}
}
