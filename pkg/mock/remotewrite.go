// Code generated by MockGen. DO NOT EDIT.
// Source: storage/remote/queue_manager.go
//
// Generated by this command:
//
//	mockgen -source storage/remote/queue_manager.go -destination mock.go -package mock WriteClient
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockWriteClient is a mock of WriteClient interface.
type MockWriteClient struct {
	ctrl     *gomock.Controller
	recorder *MockWriteClientMockRecorder
}

// MockWriteClientMockRecorder is the mock recorder for MockWriteClient.
type MockWriteClientMockRecorder struct {
	mock *MockWriteClient
}

// NewMockWriteClient creates a new mock instance.
func NewMockWriteClient(ctrl *gomock.Controller) *MockWriteClient {
	mock := &MockWriteClient{ctrl: ctrl}
	mock.recorder = &MockWriteClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriteClient) EXPECT() *MockWriteClientMockRecorder {
	return m.recorder
}

// Endpoint mocks base method.
func (m *MockWriteClient) Endpoint() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoint")
	ret0, _ := ret[0].(string)
	return ret0
}

// Endpoint indicates an expected call of Endpoint.
func (mr *MockWriteClientMockRecorder) Endpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoint", reflect.TypeOf((*MockWriteClient)(nil).Endpoint))
}

// Name mocks base method.
func (m *MockWriteClient) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockWriteClientMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockWriteClient)(nil).Name))
}

// Store mocks base method.
func (m *MockWriteClient) Store(arg0 context.Context, arg1 []byte, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockWriteClientMockRecorder) Store(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockWriteClient)(nil).Store), arg0, arg1, arg2)
}
