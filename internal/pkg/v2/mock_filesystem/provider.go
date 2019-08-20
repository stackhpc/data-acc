// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/v2/filesystem/provider.go

// Package mock_filesystem is a generated GoMock package.
package mock_filesystem

import (
	datamodel "github.com/RSE-Cambridge/data-acc/internal/pkg/v2/datamodel"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockProvider) Create(session datamodel.Session) (datamodel.FilesystemStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", session)
	ret0, _ := ret[0].(datamodel.FilesystemStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockProviderMockRecorder) Create(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProvider)(nil).Create), session)
}

// Delete mocks base method
func (m *MockProvider) Delete(session datamodel.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockProviderMockRecorder) Delete(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProvider)(nil).Delete), session)
}

// DataCopyIn mocks base method
func (m *MockProvider) DataCopyIn(session datamodel.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataCopyIn", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// DataCopyIn indicates an expected call of DataCopyIn
func (mr *MockProviderMockRecorder) DataCopyIn(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataCopyIn", reflect.TypeOf((*MockProvider)(nil).DataCopyIn), session)
}

// DataCopyOut mocks base method
func (m *MockProvider) DataCopyOut(session datamodel.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataCopyOut", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// DataCopyOut indicates an expected call of DataCopyOut
func (mr *MockProviderMockRecorder) DataCopyOut(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataCopyOut", reflect.TypeOf((*MockProvider)(nil).DataCopyOut), session)
}

// Mount mocks base method
func (m *MockProvider) Mount(session datamodel.Session, attachments datamodel.AttachmentSessionStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mount", session, attachments)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mount indicates an expected call of Mount
func (mr *MockProviderMockRecorder) Mount(session, attachments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mount", reflect.TypeOf((*MockProvider)(nil).Mount), session, attachments)
}

// Unmount mocks base method
func (m *MockProvider) Unmount(session datamodel.Session, attachments datamodel.AttachmentSessionStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmount", session, attachments)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmount indicates an expected call of Unmount
func (mr *MockProviderMockRecorder) Unmount(session, attachments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmount", reflect.TypeOf((*MockProvider)(nil).Unmount), session, attachments)
}
