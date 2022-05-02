// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/database (interfaces: Gateway)

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	api "github.com/Azure/ARO-RP/pkg/api"
	cosmosdb "github.com/Azure/ARO-RP/pkg/database/cosmosdb"
)

// MockGateway is a mock of Gateway interface.
type MockGateway struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayMockRecorder
}

// MockGatewayMockRecorder is the mock recorder for MockGateway.
type MockGatewayMockRecorder struct {
	mock *MockGateway
}

// NewMockGateway creates a new mock instance.
func NewMockGateway(ctrl *gomock.Controller) *MockGateway {
	mock := &MockGateway{ctrl: ctrl}
	mock.recorder = &MockGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGateway) EXPECT() *MockGatewayMockRecorder {
	return m.recorder
}

// ChangeFeed mocks base method.
func (m *MockGateway) ChangeFeed() cosmosdb.GatewayDocumentIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeFeed")
	ret0, _ := ret[0].(cosmosdb.GatewayDocumentIterator)
	return ret0
}

// ChangeFeed indicates an expected call of ChangeFeed.
func (mr *MockGatewayMockRecorder) ChangeFeed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeFeed", reflect.TypeOf((*MockGateway)(nil).ChangeFeed))
}

// Create mocks base method.
func (m *MockGateway) Create(arg0 context.Context, arg1 *api.GatewayDocument) (*api.GatewayDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*api.GatewayDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockGatewayMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGateway)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockGateway) Delete(arg0 context.Context, arg1 *api.GatewayDocument) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockGatewayMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGateway)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockGateway) Get(arg0 context.Context, arg1 string) (*api.GatewayDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*api.GatewayDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGatewayMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGateway)(nil).Get), arg0, arg1)
}

// Patch mocks base method.
func (m *MockGateway) Patch(arg0 context.Context, arg1 string, arg2 func(*api.GatewayDocument) error) (*api.GatewayDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.GatewayDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockGatewayMockRecorder) Patch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockGateway)(nil).Patch), arg0, arg1, arg2)
}