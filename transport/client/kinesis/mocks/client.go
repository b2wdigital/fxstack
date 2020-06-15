// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import kinesis "github.com/aws/aws-sdk-go-v2/service/kinesis"
import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// BulkPublish provides a mock function with given fields: ctx, messages, resource
func (_m *Client) BulkPublish(ctx context.Context, messages []kinesis.PutRecordsRequestEntry, resource string) error {
	ret := _m.Called(ctx, messages, resource)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []kinesis.PutRecordsRequestEntry, string) error); ok {
		r0 = rf(ctx, messages, resource)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: ctx, input
func (_m *Client) Publish(ctx context.Context, input *kinesis.PutRecordInput) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinesis.PutRecordInput) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
