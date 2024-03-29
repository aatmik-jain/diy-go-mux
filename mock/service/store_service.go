// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/aatmikjain/Documents/go-mux/service/store_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	model "go-mux/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStoreService is a mock of StoreService interface.
type MockStoreService struct {
	ctrl     *gomock.Controller
	recorder *MockStoreServiceMockRecorder
}

// MockStoreServiceMockRecorder is the mock recorder for MockStoreService.
type MockStoreServiceMockRecorder struct {
	mock *MockStoreService
}

// NewMockStoreService creates a new mock instance.
func NewMockStoreService(ctrl *gomock.Controller) *MockStoreService {
	mock := &MockStoreService{ctrl: ctrl}
	mock.recorder = &MockStoreServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreService) EXPECT() *MockStoreServiceMockRecorder {
	return m.recorder
}

// AddProducts mocks base method.
func (m *MockStoreService) AddProducts(storeID int, pIds []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProducts", storeID, pIds)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProducts indicates an expected call of AddProducts.
func (mr *MockStoreServiceMockRecorder) AddProducts(storeID, pIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProducts", reflect.TypeOf((*MockStoreService)(nil).AddProducts), storeID, pIds)
}

// GetProducts mocks base method.
func (m *MockStoreService) GetProducts(storeID int) ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", storeID)
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockStoreServiceMockRecorder) GetProducts(storeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockStoreService)(nil).GetProducts), storeID)
}
