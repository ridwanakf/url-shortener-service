// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ridwanakf/url-shortener-service/internal (interfaces: RedisJson)

// Package redisjson is a generated GoMock package.
package redisjson

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRedisJson is a mock of RedisJson interface
type MockRedisJson struct {
	ctrl     *gomock.Controller
	recorder *MockRedisJsonMockRecorder
}

// MockRedisJsonMockRecorder is the mock recorder for MockRedisJson
type MockRedisJsonMockRecorder struct {
	mock *MockRedisJson
}

// NewMockRedisJson creates a new mock instance
func NewMockRedisJson(ctrl *gomock.Controller) *MockRedisJson {
	mock := &MockRedisJson{ctrl: ctrl}
	mock.recorder = &MockRedisJsonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRedisJson) EXPECT() *MockRedisJsonMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockRedisJson) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockRedisJsonMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRedisJson)(nil).Close))
}

// Decr mocks base method
func (m *MockRedisJson) Decr(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decr", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decr indicates an expected call of Decr
func (mr *MockRedisJsonMockRecorder) Decr(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decr", reflect.TypeOf((*MockRedisJson)(nil).Decr), arg0)
}

// Del mocks base method
func (m *MockRedisJson) Del(arg0 ...interface{}) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Del indicates an expected call of Del
func (mr *MockRedisJsonMockRecorder) Del(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockRedisJson)(nil).Del), arg0...)
}

// Exists mocks base method
func (m *MockRedisJson) Exists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockRedisJsonMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockRedisJson)(nil).Exists), arg0)
}

// Expire mocks base method
func (m *MockRedisJson) Expire(arg0 string, arg1 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expire", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Expire indicates an expected call of Expire
func (mr *MockRedisJsonMockRecorder) Expire(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expire", reflect.TypeOf((*MockRedisJson)(nil).Expire), arg0, arg1)
}

// ExpireAt mocks base method
func (m *MockRedisJson) ExpireAt(arg0 string, arg1 int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireAt", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExpireAt indicates an expected call of ExpireAt
func (mr *MockRedisJsonMockRecorder) ExpireAt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireAt", reflect.TypeOf((*MockRedisJson)(nil).ExpireAt), arg0, arg1)
}

// Get mocks base method
func (m *MockRedisJson) Get(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRedisJsonMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisJson)(nil).Get), arg0)
}

// GetUnmarshalled mocks base method
func (m *MockRedisJson) GetUnmarshalled(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnmarshalled", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUnmarshalled indicates an expected call of GetUnmarshalled
func (mr *MockRedisJsonMockRecorder) GetUnmarshalled(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnmarshalled", reflect.TypeOf((*MockRedisJson)(nil).GetUnmarshalled), arg0, arg1)
}

// HDel mocks base method
func (m *MockRedisJson) HDel(arg0 string, arg1 ...string) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HDel", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HDel indicates an expected call of HDel
func (mr *MockRedisJsonMockRecorder) HDel(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HDel", reflect.TypeOf((*MockRedisJson)(nil).HDel), varargs...)
}

// HExists mocks base method
func (m *MockRedisJson) HExists(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HExists indicates an expected call of HExists
func (mr *MockRedisJsonMockRecorder) HExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HExists", reflect.TypeOf((*MockRedisJson)(nil).HExists), arg0, arg1)
}

// HGet mocks base method
func (m *MockRedisJson) HGet(arg0, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGet", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HGet indicates an expected call of HGet
func (mr *MockRedisJsonMockRecorder) HGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGet", reflect.TypeOf((*MockRedisJson)(nil).HGet), arg0, arg1)
}

// HGetAll mocks base method
func (m *MockRedisJson) HGetAll(arg0 string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGetAll", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HGetAll indicates an expected call of HGetAll
func (mr *MockRedisJsonMockRecorder) HGetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGetAll", reflect.TypeOf((*MockRedisJson)(nil).HGetAll), arg0)
}

// HKeys mocks base method
func (m *MockRedisJson) HKeys(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HKeys", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HKeys indicates an expected call of HKeys
func (mr *MockRedisJsonMockRecorder) HKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HKeys", reflect.TypeOf((*MockRedisJson)(nil).HKeys), arg0)
}

// HMGet mocks base method
func (m *MockRedisJson) HMGet(arg0 string, arg1 ...string) ([][]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HMGet", varargs...)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HMGet indicates an expected call of HMGet
func (mr *MockRedisJsonMockRecorder) HMGet(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMGet", reflect.TypeOf((*MockRedisJson)(nil).HMGet), varargs...)
}

// HSet mocks base method
func (m *MockRedisJson) HSet(arg0, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HSet indicates an expected call of HSet
func (mr *MockRedisJsonMockRecorder) HSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSet", reflect.TypeOf((*MockRedisJson)(nil).HSet), arg0, arg1, arg2)
}

// Incr mocks base method
func (m *MockRedisJson) Incr(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Incr", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Incr indicates an expected call of Incr
func (mr *MockRedisJsonMockRecorder) Incr(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockRedisJson)(nil).Incr), arg0)
}

// Set mocks base method
func (m *MockRedisJson) Set(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockRedisJsonMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisJson)(nil).Set), arg0, arg1)
}

// Setex mocks base method
func (m *MockRedisJson) Setex(arg0 string, arg1 int, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setex", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Setex indicates an expected call of Setex
func (mr *MockRedisJsonMockRecorder) Setex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setex", reflect.TypeOf((*MockRedisJson)(nil).Setex), arg0, arg1, arg2)
}

// SetexMarshalled mocks base method
func (m *MockRedisJson) SetexMarshalled(arg0 string, arg1 int, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetexMarshalled", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetexMarshalled indicates an expected call of SetexMarshalled
func (mr *MockRedisJsonMockRecorder) SetexMarshalled(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetexMarshalled", reflect.TypeOf((*MockRedisJson)(nil).SetexMarshalled), arg0, arg1, arg2)
}

// Setnx mocks base method
func (m *MockRedisJson) Setnx(arg0 string, arg1 int, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setnx", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Setnx indicates an expected call of Setnx
func (mr *MockRedisJsonMockRecorder) Setnx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setnx", reflect.TypeOf((*MockRedisJson)(nil).Setnx), arg0, arg1, arg2)
}

// TTL mocks base method
func (m *MockRedisJson) TTL(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TTL", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TTL indicates an expected call of TTL
func (mr *MockRedisJsonMockRecorder) TTL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TTL", reflect.TypeOf((*MockRedisJson)(nil).TTL), arg0)
}
