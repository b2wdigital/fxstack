package kinesis

import (
	"reflect"
	"testing"

	"github.com/b2wdigital/fxstack/domain/repository"
	"github.com/b2wdigital/fxstack/transport/client/kinesis/mocks"
	giconfig "github.com/b2wdigital/goignite/config"
	logrus "github.com/b2wdigital/goignite/log/logrus/v1"
	"github.com/stretchr/testify/suite"
)

type EventSuite struct {
	suite.Suite
}

func (s *EventSuite) SetupSuite() {
	giconfig.Load()
	logrus.NewLogger()
}

func (s *EventSuite) TestNewEvent() {

	client := new(mocks.Client)
	options, _ := DefaultOptions()

	type args struct {
		client  *mocks.Client
		options *Options
	}
	tests := []struct {
		name string
		args args
		want repository.Event
	}{
		{
			name: "Success",
			args: args{
				client:  client,
				options: options,
			},
			want: NewEvent(client, options),
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewEvent(tt.args.client, tt.args.options)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewEvent() = %v, want %v", got, tt.want)
		})
	}
}

func TestEventSuite(t *testing.T) {
	suite.Run(t, new(EventSuite))
}
