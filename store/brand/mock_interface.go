// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package brand is a generated GoMock package.
package brand

import (
	model "Catlog/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// GetById mocks base method.
func (m *MockStore) GetById(id int) (model.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(model.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockStoreMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockStore)(nil).GetById), id)
}

// GetByName mocks base method.
func (m *MockStore) GetByName(name string) (model.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", name)
	ret0, _ := ret[0].(model.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockStoreMockRecorder) GetByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockStore)(nil).GetByName), name)
}

// InsertBrand mocks base method.
func (m *MockStore) InsertBrand(p model.Brand) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBrand", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertBrand indicates an expected call of InsertBrand.
func (mr *MockStoreMockRecorder) InsertBrand(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBrand", reflect.TypeOf((*MockStore)(nil).InsertBrand), p)
}
