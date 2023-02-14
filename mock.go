// Code generated by MockGen. DO NOT EDIT.
// Source: ./codec.go

// Package codec is a generated GoMock package.
package codec

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCodec is a mock of Codec interface.
type MockCodec struct {
	ctrl     *gomock.Controller
	recorder *MockCodecMockRecorder
}

// MockCodecMockRecorder is the mock recorder for MockCodec.
type MockCodecMockRecorder struct {
	mock *MockCodec
}

// NewMockCodec creates a new mock instance.
func NewMockCodec(ctrl *gomock.Controller) *MockCodec {
	mock := &MockCodec{ctrl: ctrl}
	mock.recorder = &MockCodecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodec) EXPECT() *MockCodecMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockCodec) Decode(r io.Reader, dst interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", r, dst)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockCodecMockRecorder) Decode(r, dst interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockCodec)(nil).Decode), r, dst)
}

// Encode mocks base method.
func (m *MockCodec) Encode(data interface{}) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", data)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode.
func (mr *MockCodecMockRecorder) Encode(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockCodec)(nil).Encode), data)
}
