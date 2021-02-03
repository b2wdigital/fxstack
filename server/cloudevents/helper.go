package cloudevents

import (
	"context"
	"log"

	"github.com/b2wdigital/fxstack/cloudevents"
	ce "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
)

type Helper struct {
	client  client.Client
	handler cloudevents.Handler
	ctx     context.Context
}

func NewHelper(ctx context.Context, handler cloudevents.Handler) (*Helper, error) {

	c, err := ce.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	return &Helper{
		ctx:     ctx,
		client:  c,
		handler: NewHandler(handler),
	}, nil
}

func (h *Helper) Serve() {
	log.Fatal(h.client.StartReceiver(h.ctx, h.handler.Handle))
}
