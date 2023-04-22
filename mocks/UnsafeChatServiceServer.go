// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeChatServiceServer is an autogenerated mock type for the UnsafeChatServiceServer type
type UnsafeChatServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedChatServiceServer provides a mock function with given fields:
func (_m *UnsafeChatServiceServer) mustEmbedUnimplementedChatServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewUnsafeChatServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeChatServiceServer creates a new instance of UnsafeChatServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeChatServiceServer(t mockConstructorTestingTNewUnsafeChatServiceServer) *UnsafeChatServiceServer {
	mock := &UnsafeChatServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
