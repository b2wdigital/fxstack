package lambda

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/b2wdigital/fxstack/cloudevents"
	"github.com/b2wdigital/fxstack/cloudevents/middleware"
	"github.com/b2wdigital/fxstack/cloudevents/mocks"
	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HandlerSuite struct {
	suite.Suite
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (s *HandlerSuite) SetupSuite() {
	giconfig.Load()
	gilogrus.NewLogger()
}

func (s *HandlerSuite) TestHandler_Handle() {

	handler := new(mocks.Handler)

	lc := new(lambdacontext.LambdaContext)
	ctx := lambdacontext.NewContext(context.Background(), lc)

	var kinesisEvent1 Event
	b, _ := ioutil.ReadFile("testdata/kinesis_success.json")
	json.Unmarshal(b, &kinesisEvent1)

	var middlewares []cloudevents.Middleware

	middlewares = append(middlewares, middleware.NewLog())

	options, _ := DefaultOptions()

	type fields struct {
		handler     *mocks.Handler
		middlewares []cloudevents.Middleware
		options     *Options
	}

	type args struct {
		ctx   context.Context
		event Event
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(handler *mocks.Handler)
	}{
		{
			name: "on kinesis success event",
			fields: fields{
				handler:     handler,
				middlewares: middlewares,
				options:     options,
			},
			args: args{
				ctx:   ctx,
				event: kinesisEvent1,
			},
			wantErr: false,
			mock: func(handler *mocks.Handler) {

				e := v2.NewEvent()
				e.SetSubject("changeme")
				e.SetSource("changeme")
				e.SetType("changeme")
				e.SetData("", "changeme")

				handler.On("Handle", mock.Anything, mock.Anything).Times(1).
					Return(&e, nil)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {

			tt.mock(tt.fields.handler)

			h := NewHandler(tt.fields.handler, tt.fields.middlewares, tt.fields.options)

			err := h.Handle(tt.args.ctx, tt.args.event)
			if err != nil {
				gilog.Error(err)
			}

			s.Assert().True((err != nil) == tt.wantErr, "Handle() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}

func (s *HandlerSuite) TestNewHandler() {

	type args struct {
		handler     cloudevents.Handler
		middlewares []cloudevents.Middleware
		options     *Options
	}

	handler := new(mocks.Handler)
	options, _ := DefaultOptions()

	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "success",
			args: args{
				handler:     handler,
				middlewares: nil,
				options:     options,
			},
			want: NewHandler(handler, nil, options),
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {

			got := NewHandler(tt.args.handler, tt.args.middlewares, tt.args.options)

			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewHandler() = %v, want %v")

		})
	}
}
