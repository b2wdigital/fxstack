package lambda

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
	v2 "github.com/cloudevents/sdk-go/v2"
	"golang.org/x/sync/errgroup"
)

func fromSQS(parentCtx context.Context, event Event) []*cloudevents.InOut {

	logger := gilog.FromContext(parentCtx)
	logger.Info("receiving SQS event")

	lc, _ := lambdacontext.FromContext(parentCtx)

	var inouts []*cloudevents.InOut

	g, gctx := errgroup.WithContext(parentCtx)

	for _, record := range event.Records {

		record := record

		g.Go(func() error {

			var err error

			in := v2.NewEvent()

			if err = json.Unmarshal([]byte(record.SQS.Body), &in); err != nil {
				var data interface{}

				if err = json.Unmarshal([]byte(record.SQS.Body), &data); err != nil {
					err = errors.NewNotValid(err, "could not decode SQS record")
				} else {
					if err = in.SetData("", data); err != nil {
						err = errors.NewNotValid(err, "could not decode SQS record")
					}
				}
			}

			in.SetType(record.EventSource)

			if in.ID() == "" {
				in.SetID(record.SQS.MessageId)
			}

			in.SetSource(record.EventSource)

			in.SetExtension("awsRequestID", lc.AwsRequestID)
			in.SetExtension("invokedFunctionArn", lc.InvokedFunctionArn)

			inouts = append(inouts, &cloudevents.InOut{
				In:  &in,
				Err: err,
			})

			return nil

		})

	}

	if err := g.Wait(); err == nil {
		logger.Debug("all events converted")
	}

	gctx.Done()

	return inouts

}
