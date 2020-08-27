package lambda

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
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

			j, _ := json.Marshal(record)
			logger.Debug(string(j))

			var err error

			in := v2.NewEvent()

			if er := json.Unmarshal([]byte(record.Body), &in); er != nil {

				var data interface{}

				snsEvent := events.SNSEntity{}

				if err = json.Unmarshal([]byte(record.Body), &snsEvent); err == nil {

					if snsEvent.Message != "" {

						if er := json.Unmarshal([]byte(snsEvent.Message), &data); er != nil {
							err = errors.NewNotValid(er, "could not decode SNS message from SQS record")
						}

					} else {

						if er := json.Unmarshal([]byte(record.Body), &data); er != nil {
							err = errors.NewNotValid(er, "could not decode SQS record")
						}

					}

					if er := in.SetData("", data); er != nil {
						err = errors.NewNotValid(er, "could not set data in event")
					}

				}

			}

			in.SetType(record.EventSource)

			if in.ID() == "" {
				in.SetID(record.MessageId)
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
