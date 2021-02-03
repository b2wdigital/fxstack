package cloudevents

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/fxstack/util"
	gilog "github.com/b2wdigital/goignite/log"
	v2 "github.com/cloudevents/sdk-go/v2"
)

type Handler struct {
	handler cloudevents.Handler
}

func NewHandler(handler cloudevents.Handler) cloudevents.Handler {
	return &Handler{handler: handler}
}

func (h *Handler) Handle(parentCtx context.Context, in v2.Event) (out *v2.Event, err error) {

	if util.StringSliceContains([]string{"com.amazon.sqs.message", "aws.sqs.message"}, in.Type()) {
		h.fromSQS(parentCtx, &in)
	}

	return h.handler.Handle(parentCtx, in)
}

func (h *Handler) fromSQS(parentCtx context.Context, in *v2.Event) {

	logger := gilog.FromContext(parentCtx)

	var err error
	var sqsMessage events.SQSMessage

	err = json.Unmarshal(in.Data(), &sqsMessage)
	if err != nil {
		logger.Error(err)
	}

	var snsEntity events.SNSEntity

	err = json.Unmarshal([]byte(sqsMessage.Body), &snsEntity)
	if err != nil {
		logger.Error(err)
	}

	var data []byte

	if snsEntity.Message != "" {
		data = []byte(snsEntity.Message)
	} else {
		data = []byte(sqsMessage.Body)
	}

	ctype := v2.TextPlain
	var js map[string]interface{}

	if json.Unmarshal(data, &js) == nil {
		ctype = v2.ApplicationJSON
	}

	err = in.SetData(ctype, data)
	if err != nil {
		logger.Error(err)
	}

}
