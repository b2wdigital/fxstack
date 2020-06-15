package cloudevents

import (
	"context"

	v2 "github.com/cloudevents/sdk-go/v2"
)

type InOut struct {
	In      *v2.Event
	Out     *v2.Event
	Err     error
	Context context.Context
}
