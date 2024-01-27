// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/store.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	store "github.com/sebasttiano/alice-skill/internal/store"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// FindRecipient mocks base method.
func (m *MockStore) FindRecipient(ctx context.Context, username string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRecipient", ctx, username)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRecipient indicates an expected call of FindRecipient.
func (mr *MockStoreMockRecorder) FindRecipient(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRecipient", reflect.TypeOf((*MockStore)(nil).FindRecipient), ctx, username)
}

// GetMessage mocks base method.
func (m *MockStore) GetMessage(ctx context.Context, id int64) (*store.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage", ctx, id)
	ret0, _ := ret[0].(*store.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessage indicates an expected call of GetMessage.
func (mr *MockStoreMockRecorder) GetMessage(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockStore)(nil).GetMessage), ctx, id)
}

// ListMessages mocks base method.
func (m *MockStore) ListMessages(ctx context.Context, userID string) ([]store.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMessages", ctx, userID)
	ret0, _ := ret[0].([]store.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMessages indicates an expected call of ListMessages.
func (mr *MockStoreMockRecorder) ListMessages(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMessages", reflect.TypeOf((*MockStore)(nil).ListMessages), ctx, userID)
}

// SaveMessage mocks base method.
func (m *MockStore) SaveMessage(ctx context.Context, userID string, msg store.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveMessage", ctx, userID, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveMessage indicates an expected call of SaveMessage.
func (mr *MockStoreMockRecorder) SaveMessage(ctx, userID, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveMessage", reflect.TypeOf((*MockStore)(nil).SaveMessage), ctx, userID, msg)
}
