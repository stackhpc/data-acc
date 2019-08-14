// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/v2/registry/brick.go

// Package mock_registry is a generated GoMock package.
package mock_registry

import (
	context "context"
	datamodel "github.com/RSE-Cambridge/data-acc/internal/pkg/v2/datamodel"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBrickRegistry is a mock of BrickRegistry interface
type MockBrickRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockBrickRegistryMockRecorder
}

// MockBrickRegistryMockRecorder is the mock recorder for MockBrickRegistry
type MockBrickRegistryMockRecorder struct {
	mock *MockBrickRegistry
}

// NewMockBrickRegistry creates a new mock instance
func NewMockBrickRegistry(ctrl *gomock.Controller) *MockBrickRegistry {
	mock := &MockBrickRegistry{ctrl: ctrl}
	mock.recorder = &MockBrickRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBrickRegistry) EXPECT() *MockBrickRegistryMockRecorder {
	return m.recorder
}

// UpdateBrickHost mocks base method
func (m *MockBrickRegistry) UpdateBrickHost(brickHostInfo datamodel.BrickHost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBrickHost", brickHostInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBrickHost indicates an expected call of UpdateBrickHost
func (mr *MockBrickRegistryMockRecorder) UpdateBrickHost(brickHostInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBrickHost", reflect.TypeOf((*MockBrickRegistry)(nil).UpdateBrickHost), brickHostInfo)
}

// GetSessionActions mocks base method
func (m *MockBrickRegistry) GetSessionActions(ctxt context.Context, brickHostName datamodel.BrickHostName) (<-chan datamodel.SessionAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionActions", ctxt, brickHostName)
	ret0, _ := ret[0].(<-chan datamodel.SessionAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionActions indicates an expected call of GetSessionActions
func (mr *MockBrickRegistryMockRecorder) GetSessionActions(ctxt, brickHostName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionActions", reflect.TypeOf((*MockBrickRegistry)(nil).GetSessionActions), ctxt, brickHostName)
}

// KeepAliveHost mocks base method
func (m *MockBrickRegistry) KeepAliveHost(brickHostName datamodel.BrickHostName) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeepAliveHost", brickHostName)
	ret0, _ := ret[0].(error)
	return ret0
}

// KeepAliveHost indicates an expected call of KeepAliveHost
func (mr *MockBrickRegistryMockRecorder) KeepAliveHost(brickHostName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeepAliveHost", reflect.TypeOf((*MockBrickRegistry)(nil).KeepAliveHost), brickHostName)
}

// IsBrickHostAlive mocks base method
func (m *MockBrickRegistry) IsBrickHostAlive(brickHostName datamodel.BrickHostName) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBrickHostAlive", brickHostName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsBrickHostAlive indicates an expected call of IsBrickHostAlive
func (mr *MockBrickRegistryMockRecorder) IsBrickHostAlive(brickHostName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBrickHostAlive", reflect.TypeOf((*MockBrickRegistry)(nil).IsBrickHostAlive), brickHostName)
}