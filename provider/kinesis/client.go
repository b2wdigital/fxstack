package kinesis

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
	v2 "github.com/cloudevents/sdk-go/v2"

	k "github.com/b2wdigital/fxstack/transport/client/kinesis"
)

type Client struct {
	client  k.Client
	options *Options
}

func NewClient(c k.Client, options *Options) *Client {
	return &Client{client: c, options: options}
}

func (p *Client) Publish(ctx context.Context, outs []*v2.Event) error {

	logger := gilog.FromContext(ctx).WithTypeOf(*p)

	logger.Info("publishing to kinesis")

	if len(outs) > 1 {

		return p.multi(ctx, outs)

	} else if len(outs) == 1 {

		return p.single(ctx, outs)

	}

	logger.Warnf("no messages were reported for posting")

	return nil
}

func (p *Client) multi(ctx context.Context, outs []*v2.Event) (err error) {

	logger := gilog.FromContext(ctx).WithTypeOf(*p)

	bulks := make(map[string][]*types.PutRecordsRequestEntry)

	for _, out := range outs {

		var rawMessage []byte

		rawMessage, err = p.rawMessage(out)
		if err != nil {
			return errors.Wrap(err, errors.Internalf("error on marshal. %s", err.Error()))
		}

		var partitionKey string

		partitionKey, err = p.partitionKey(out)
		if err != nil {
			return err
		}

		entry := &types.PutRecordsRequestEntry{
			Data:         rawMessage,
			PartitionKey: aws.String(partitionKey),
		}

		logger.WithField("partitionKey", partitionKey).
			WithField("subject", out.Subject()).
			WithField("id", out.ID()).
			Info(string(rawMessage))

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

func (p *Client) single(ctx context.Context, outs []*v2.Event) (err error) {

	logger := gilog.FromContext(ctx).WithTypeOf(*p)

	out := outs[0]

	var rawMessage []byte

	rawMessage, err = p.rawMessage(out)
	if err != nil {
		return errors.Wrap(err, errors.Internalf("error on marshal. %s", err.Error()))
	}

	var partitionKey string

	partitionKey, err = p.partitionKey(out)
	if err != nil {
		return err
	}

	input := &kinesis.PutRecordInput{
		Data:         rawMessage,
		PartitionKey: aws.String(partitionKey),
		StreamName:   aws.String(out.Subject()),
	}

	logger.WithField("partitionKey", partitionKey).
		WithField("subject", out.Subject()).
		WithField("id", out.ID()).
		Info(string(rawMessage))

	err = p.client.Publish(ctx, input)
	if err != nil {
		return errors.NewInternal(err, "could not be published in kinesis")
	}

	return nil
}

func (p *Client) rawMessage(out *v2.Event) (rawMessage []byte, err error) {
	exts := out.Extensions()

	source, ok := exts["target"]

	if ok {

		s := source.(string)

		if s == "data" {
			var data interface{}

			err = out.DataAs(&data)
			if err != nil {
				return nil, errors.Wrap(err, errors.Internalf("error on data as. %s", err.Error()))
			}

			rawMessage, err = json.Marshal(data)

		} else {
			rawMessage, err = cloudevents.JSONBytes(*out)
		}
	} else {
		rawMessage, err = cloudevents.JSONBytes(*out)
	}

	return rawMessage, err
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
