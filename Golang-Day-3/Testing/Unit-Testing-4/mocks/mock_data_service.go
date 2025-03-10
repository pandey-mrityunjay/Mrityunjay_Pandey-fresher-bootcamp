// Code generated by MockGen. DO NOT EDIT.
// Source: services/data_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDataService is a mock of DataService interface.
type MockDataService struct {
	ctrl     *gomock.Controller
	recorder *MockDataServiceMockRecorder
}

// MockDataServiceMockRecorder is the mock recorder for MockDataService.
type MockDataServiceMockRecorder struct {
	mock *MockDataService
}

// NewMockDataService creates a new mock instance.
func NewMockDataService(ctrl *gomock.Controller) *MockDataService {
	mock := &MockDataService{ctrl: ctrl}
	mock.recorder = &MockDataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataService) EXPECT() *MockDataServiceMockRecorder {
	return m.recorder
}

// GetData mocks base method.
func (m *MockDataService) GetData(id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *MockDataServiceMockRecorder) GetData(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockDataService)(nil).GetData), id)
}
