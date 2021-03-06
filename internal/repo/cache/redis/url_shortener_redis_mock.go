// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ridwanakf/url-shortener-service/internal (interfaces: ShortenerCacheRepo)

// Package redis is a generated GoMock package.
package redis

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/ridwanakf/url-shortener-service/internal/entity"
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
func (m *MockShortenerCacheRepo) DeleteURL(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteURL", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteURL indicates an expected call of DeleteURL
func (mr *MockShortenerCacheRepoMockRecorder) DeleteURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteURL", reflect.TypeOf((*MockShortenerCacheRepo)(nil).DeleteURL), arg0)
}

// GetAllURL mocks base method
func (m *MockShortenerCacheRepo) GetAllURL(arg0 string) ([]entity.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllURL", arg0)
	ret0, _ := ret[0].([]entity.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllURL indicates an expected call of GetAllURL
func (mr *MockShortenerCacheRepoMockRecorder) GetAllURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllURL", reflect.TypeOf((*MockShortenerCacheRepo)(nil).GetAllURL), arg0)
}

// GetURL mocks base method
func (m *MockShortenerCacheRepo) GetURL(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURL indicates an expected call of GetURL
func (mr *MockShortenerCacheRepoMockRecorder) GetURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockShortenerCacheRepo)(nil).GetURL), arg0)
}

// HasShortURLExpired mocks base method
func (m *MockShortenerCacheRepo) HasShortURLExpired(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasShortURLExpired", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasShortURLExpired indicates an expected call of HasShortURLExpired
func (mr *MockShortenerCacheRepoMockRecorder) HasShortURLExpired(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasShortURLExpired", reflect.TypeOf((*MockShortenerCacheRepo)(nil).HasShortURLExpired), arg0)
}

// IsCollectionURLExist mocks base method
func (m *MockShortenerCacheRepo) IsCollectionURLExist(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCollectionURLExist", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsCollectionURLExist indicates an expected call of IsCollectionURLExist
func (mr *MockShortenerCacheRepoMockRecorder) IsCollectionURLExist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCollectionURLExist", reflect.TypeOf((*MockShortenerCacheRepo)(nil).IsCollectionURLExist), arg0)
}

// IsSingleURLExist mocks base method
func (m *MockShortenerCacheRepo) IsSingleURLExist(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSingleURLExist", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSingleURLExist indicates an expected call of IsSingleURLExist
func (mr *MockShortenerCacheRepoMockRecorder) IsSingleURLExist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSingleURLExist", reflect.TypeOf((*MockShortenerCacheRepo)(nil).IsSingleURLExist), arg0)
}

// SetURL mocks base method
func (m *MockShortenerCacheRepo) SetURL(arg0 entity.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetURL", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetURL indicates an expected call of SetURL
func (mr *MockShortenerCacheRepoMockRecorder) SetURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetURL", reflect.TypeOf((*MockShortenerCacheRepo)(nil).SetURL), arg0)
}
