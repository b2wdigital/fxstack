package lambda

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/b2wdigital/goignite/errors"
	gilog "github.com/b2wdigital/goignite/log"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/b2wdigital/fxstack/cloudevents"
)

func fromKinesis(parentCtx context.Context, event Event) []*cloudevents.InOut {

	logger := gilog.FromContext(parentCtx)
	logger.Info("receiving Kinesis event")

	lc, _ := lambdacontext.FromContext(parentCtx)

	var inouts []*cloudevents.InOut

	j, _ := json.Marshal(event)
	gilog.Debug(string(j))

	for _, record := range event.Records {
		var err error
		in := v2.NewEvent()

		if err = json.Unmarshal(record.Kinesis.Data, &in); err != nil {
			var data interface{}

			if err = json.Unmarshal(record.Kinesis.Data, &data); err != nil {
				err = errors.NewNotValid(err, "could not decode kinesis record")
			} else {
				if err = in.SetData("", data); err != nil {
					err = errors.NewNotValid(err, "could not decode kinesis record")
				}
			}
		}

		in.SetType(record.EventName)

		if in.ID() == "" {
			in.SetID(record.EventID)
		}

		in.SetSource(record.EventSource)

		in.SetExtension("awsRequestID", lc.AwsRequestID)
		in.SetExtension("invokedFunctionArn", lc.InvokedFunctionArn)

		inouts = append(inouts, &cloudevents.InOut{
			In:  &in,
			Err: err,
		})
	}

	return inouts
}
