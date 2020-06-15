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

func fromSNS(parentCtx context.Context, event Event) []*cloudevents.InOut {

	logger := gilog.FromContext(parentCtx)
	logger.Info("receiving SNS event")

	lc, _ := lambdacontext.FromContext(parentCtx)

	var inouts []*cloudevents.InOut

	g, gctx := errgroup.WithContext(parentCtx)

	for _, record := range event.Records {

		record := record

		g.Go(func() error {

			var err error

			in := v2.NewEvent()

			if err = json.Unmarshal([]byte(record.SNS.Message), &in); err != nil {

				var data interface{}

				if err = json.Unmarshal([]byte(record.SNS.Message), &data); err != nil {
					err = errors.NewNotValid(err, "could not decode SNS record")
				} else {

					err = in.SetData("", data)
					if err != nil {
						err = errors.NewNotValid(err, "could not decode SNS record")
					}
				}

			}

			in.SetType(record.SNS.Type)
			in.SetID(record.SNS.MessageID)
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
