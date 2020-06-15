package nats

import (
	"context"
	"fmt"
	"testing"

	giconfig "github.com/b2wdigital/goignite/config"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	ginats "github.com/b2wdigital/goignite/nats/v1"
	"github.com/cloudevents/sdk-go/v2"

	"github.com/nats-io/gnatsd/server"
	natsserver "github.com/nats-io/nats-server/test"

	"github.com/stretchr/testify/assert"
)

const TestPort = 8369

type Handler struct {
}

func (h *Handler) Handle(ctx context.Context, in v2.Event) (*v2.Event, error) {
	fmt.Println("hello world")
	return nil, nil
}

func runServerOnPort(port int) *server.Server {
	opts := natsserver.DefaultTestOptions
	opts.Port = port
	return runServerWithOptions(&opts)
}

func runServerWithOptions(opts *server.Options) *server.Server {
	return natsserver.RunServer(opts)
}

func TestSubscriberListenerSubscribe(t *testing.T) {

	giconfig.Load()
	gilogrus.NewLogger()

	var err error
	var options *ginats.Options

	s := runServerOnPort(TestPort)
	defer s.Shutdown()

	sUrl := fmt.Sprintf("nats://127.0.0.1:%d", TestPort)

	options, err = ginats.DefaultOptions()
	assert.Nil(t, err)

	options.Url = sUrl

	conn, err := ginats.NewConnection(context.Background(), options)
	assert.Nil(t, err)
	defer conn.Close()

	q, err := ginats.NewQueue(context.Background(), options)
	assert.Nil(t, err)

	lis := NewSubscriberListener(q, nil, "subject", "queue")
	subscribe, err := lis.Subscribe(context.Background())
	assert.Nil(t, err)

	assert.True(t, subscribe.IsValid())

	err = subscribe.Unsubscribe()
	assert.Nil(t, err)
}
