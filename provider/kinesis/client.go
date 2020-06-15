package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/b2wdigital/fxstack/cloudevents"
	k "github.com/b2wdigital/fxstack/transport/client/kinesis"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
	v2 "github.com/cloudevents/sdk-go/v2"
)

// client holds evertything needed to publish a product
type Client struct {
	client  k.Client
	options *Options
}

func NewClient(c k.Client, options *Options) *Client {
	return &Client{client: c, options: options}
}

// Publish publishes an array of products on
func (p *Client) Publish(ctx context.Context, outs []*v2.Event) error {

	logger := gilog.FromContext(ctx).WithTypeOf(*p)

	if len(outs) > 1 {

		return p.multi(ctx, outs)

	} else if len(outs) == 1 {

		return p.single(ctx, outs)

	} else {

		logger.Warnf("no messages were reported for posting")

	}

	return nil
}

func (p *Client) multi(ctx context.Context, outs []*v2.Event) error {

	bulks := make(map[string][]kinesis.PutRecordsRequestEntry)

	for _, out := range outs {

		rawMessage, err := cloudevents.JSONBytes(*out)
		if err != nil {
			return errors.Wrap(err, errors.Internalf("error on marshal. %s", err.Error()))
		}

		partitionKey, err := p.partitionKey(out)
		if err != nil {
			return err
		}

		entry := kinesis.PutRecordsRequestEntry{
			Data:         rawMessage,
			PartitionKey: aws.String(partitionKey),
		}

		bulks[out.Subject()] = append(bulks[out.Subject()], entry)
	}

	for subject, events := range bulks {
		err := p.client.BulkPublish(ctx, events, subject)
		if err != nil {
			return errors.NewInternal(err, "could not be published in kinesis")
		}
	}

	return nil
}

func (p *Client) single(ctx context.Context, outs []*v2.Event) error {

	rawMessage, err := cloudevents.JSONBytes(*outs[0])
	if err != nil {
		return errors.Wrap(err, errors.Internalf("error on marshal. %s", err.Error()))
	}

	partitionKey, err := p.partitionKey(outs[0])
	if err != nil {
		return err
	}

	input := &kinesis.PutRecordInput{
		Data:         rawMessage,
		PartitionKey: aws.String(partitionKey),
		StreamName:   aws.String(outs[0].Subject()),
	}

	err = p.client.Publish(ctx, input)
	if err != nil {
		return errors.NewInternal(err, "could not be published in kinesis")
	}

	return nil
}

func (p *Client) partitionKey(out *v2.Event) (string, error) {

	var pk string
	exts := out.Extensions()

	if group, ok := exts["group"]; ok {
		pk = group.(string)
	} else {
		pk = "unknown"
	}

	return pk, nil
}
