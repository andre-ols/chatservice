// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/andre-ols/chatservice/internal/infra/grpc/pb"
)

// ChatServiceClient is an autogenerated mock type for the ChatServiceClient type
type ChatServiceClient struct {
	mock.Mock
}

// ChatStream provides a mock function with given fields: ctx, in, opts
func (_m *ChatServiceClient) ChatStream(ctx context.Context, in *pb.ChatRequest, opts ...grpc.CallOption) (pb.ChatService_ChatStreamClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 pb.ChatService_ChatStreamClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ChatRequest, ...grpc.CallOption) (pb.ChatService_ChatStreamClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ChatRequest, ...grpc.CallOption) pb.ChatService_ChatStreamClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pb.ChatService_ChatStreamClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.ChatRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewChatServiceClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewChatServiceClient creates a new instance of ChatServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChatServiceClient(t mockConstructorTestingTNewChatServiceClient) *ChatServiceClient {
	mock := &ChatServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
