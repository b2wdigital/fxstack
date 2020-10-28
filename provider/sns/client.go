package sns

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/matryer/try"
	"golang.org/x/sync/errgroup"

	k "github.com/b2wdigital/fxstack/transport/client/sns"
)

type Client struct {
	client  k.Client
}

func NewClient(c k.Client) *Client {
	return &Client{client: c}
}

func (p *Client) Publish(ctx context.Context, events []*v2.Event) error {

	logger := gilog.FromContext(ctx).WithTypeOf(*p)

	logger.Info("publishing to sns")

	if len(events) > 0 {

		return p.send(ctx, events)

	}

	logger.Warnf("no messages were reported for posting")

	return nil
}

func (p *Client) send(parentCtx context.Context, events []*v2.Event) (err error) {

	logger := gilog.FromContext(parentCtx).WithTypeOf(*p)

	g, gctx := errgroup.WithContext(parentCtx)
	defer gctx.Done()

	for _, e := range events {

		event := e

		g.Go(func() (err error) {

			var rawMessage []byte

			rawMessage, err = p.rawMessage(event)
			if err != nil {
				return errors.Wrap(err, errors.Internalf("error on marshal. %s", err.Error()))
			}

			message := Message{
				Default: string(rawMessage),
			}
			messageBytes, _ := json.Marshal(message)
			messageStr := string(messageBytes)

			input := &sns.PublishInput{
				Message:          aws.String(messageStr),
				MessageStructure: aws.String("json"),
				TopicArn:         aws.String(event.Subject()),
			}

			logger.WithField("subject", event.Subject()).
				WithField("id", event.ID()).
				Info(string(rawMessage))

			err = try.Do(func(attempt int) (bool, error) {
				var err error
				err = p.client.Publish(gctx, input)
				if err != nil {
					return attempt < 5, errors.NewInternal(err, "could not be published in sns")
				}
				return false, nil
			})

			return err

		})


	}

	return g.Wait()
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

type Message struct {
	Default string `json:"default"`
}
