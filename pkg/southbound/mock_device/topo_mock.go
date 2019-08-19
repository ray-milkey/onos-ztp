// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onosproject/onos-topo/pkg/northbound/device (interfaces: DeviceServiceClient)

// Package mock_device is a generated GoMock package.
package mock_device

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	device "github.com/onosproject/onos-topo/pkg/northbound/device"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockDeviceServiceClient is a mock of DeviceServiceClient interface
type MockDeviceServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceServiceClientMockRecorder
}

// MockDeviceServiceClientMockRecorder is the mock recorder for MockDeviceServiceClient
type MockDeviceServiceClientMockRecorder struct {
	mock *MockDeviceServiceClient
}

// NewMockDeviceServiceClient creates a new mock instance
func NewMockDeviceServiceClient(ctrl *gomock.Controller) *MockDeviceServiceClient {
	mock := &MockDeviceServiceClient{ctrl: ctrl}
	mock.recorder = &MockDeviceServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceServiceClient) EXPECT() *MockDeviceServiceClientMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockDeviceServiceClient) Add(arg0 context.Context, arg1 *device.AddRequest, arg2 ...grpc.CallOption) (*device.AddResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Add", varargs...)
	ret0, _ := ret[0].(*device.AddResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockDeviceServiceClientMockRecorder) Add(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockDeviceServiceClient)(nil).Add), varargs...)
}

// Get mocks base method
func (m *MockDeviceServiceClient) Get(arg0 context.Context, arg1 *device.GetRequest, arg2 ...grpc.CallOption) (*device.GetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*device.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDeviceServiceClientMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeviceServiceClient)(nil).Get), varargs...)
}

// List mocks base method
func (m *MockDeviceServiceClient) List(arg0 context.Context, arg1 *device.ListRequest, arg2 ...grpc.CallOption) (device.DeviceService_ListClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(device.DeviceService_ListClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockDeviceServiceClientMockRecorder) List(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDeviceServiceClient)(nil).List), varargs...)
}

// Remove mocks base method
func (m *MockDeviceServiceClient) Remove(arg0 context.Context, arg1 *device.RemoveRequest, arg2 ...grpc.CallOption) (*device.RemoveResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Remove", varargs...)
	ret0, _ := ret[0].(*device.RemoveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Remove indicates an expected call of Remove
func (mr *MockDeviceServiceClientMockRecorder) Remove(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockDeviceServiceClient)(nil).Remove), varargs...)
}

// Update mocks base method
func (m *MockDeviceServiceClient) Update(arg0 context.Context, arg1 *device.UpdateRequest, arg2 ...grpc.CallOption) (*device.UpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*device.UpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockDeviceServiceClientMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDeviceServiceClient)(nil).Update), varargs...)
}