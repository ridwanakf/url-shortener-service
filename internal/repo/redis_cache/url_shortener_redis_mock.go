// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ridwanakf/url-shortener-service/internal (interfaces: ShortenerCacheRepo)

// Package redis_cache is a generated GoMock package.
package redis_cache

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockShortenerCacheRepo is a mock of ShortenerCacheRepo interface
type MockShortenerCacheRepo struct {
	ctrl     *gomock.Controller
	recorder *MockShortenerCacheRepoMockRecorder
}

// MockShortenerCacheRepoMockRecorder is the mock recorder for MockShortenerCacheRepo
type MockShortenerCacheRepoMockRecorder struct {
	mock *MockShortenerCacheRepo
}

// NewMockShortenerCacheRepo creates a new mock instance
func NewMockShortenerCacheRepo(ctrl *gomock.Controller) *MockShortenerCacheRepo {
	mock := &MockShortenerCacheRepo{ctrl: ctrl}
	mock.recorder = &MockShortenerCacheRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockShortenerCacheRepo) EXPECT() *MockShortenerCacheRepoMockRecorder {
	return m.recorder
}

// DeleteURL mocks base method
func (m *MockShortenerCacheRepo) DeleteURL(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteURL", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteURL indicates an expected call of DeleteURL
func (mr *MockShortenerCacheRepoMockRecorder) DeleteURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteURL", reflect.TypeOf((*MockShortenerCacheRepo)(nil).DeleteURL), arg0)
}

// GetLongURL mocks base method
func (m *MockShortenerCacheRepo) GetLongURL(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLongURL indicates an expected call of GetLongURL
func (mr *MockShortenerCacheRepoMockRecorder) GetLongURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongURL", reflect.TypeOf((*MockShortenerCacheRepo)(nil).GetLongURL), arg0)
}

// IsShortURLExist mocks base method
func (m *MockShortenerCacheRepo) IsShortURLExist(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsShortURLExist", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsShortURLExist indicates an expected call of IsShortURLExist
func (mr *MockShortenerCacheRepoMockRecorder) IsShortURLExist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsShortURLExist", reflect.TypeOf((*MockShortenerCacheRepo)(nil).IsShortURLExist), arg0)
}
