// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package service is a generated GoMock package.
package service

import (
	model "Catlog/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProduct is a mock of Product interface.
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct.
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance.
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockProduct) GetById(id int) (model.NewProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(model.NewProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockProductMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockProduct)(nil).GetById), id)
}

// InsertProductBrand mocks base method.
func (m *MockProduct) InsertProductBrand(pName, bName string) (model.NewProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProductBrand", pName, bName)
	ret0, _ := ret[0].(model.NewProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertProductBrand indicates an expected call of InsertProductBrand.
func (mr *MockProductMockRecorder) InsertProductBrand(pName, bName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProductBrand", reflect.TypeOf((*MockProduct)(nil).InsertProductBrand), pName, bName)
}
