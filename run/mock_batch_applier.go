// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/box/kube-applier/run

package run

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of BatchApplierInterface interface
type MockBatchApplierInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockBatchApplierInterfaceRecorder
}

// Recorder for MockBatchApplierInterface (not exported)
type _MockBatchApplierInterfaceRecorder struct {
	mock *MockBatchApplierInterface
}

func NewMockBatchApplierInterface(ctrl *gomock.Controller) *MockBatchApplierInterface {
	mock := &MockBatchApplierInterface{ctrl: ctrl}
	mock.recorder = &_MockBatchApplierInterfaceRecorder{mock}
	return mock
}

func (_m *MockBatchApplierInterface) EXPECT() *_MockBatchApplierInterfaceRecorder {
	return _m.recorder
}

func (_m *MockBatchApplierInterface) Apply(_param0 []string) ([]ApplyAttempt, []ApplyAttempt) {
	ret := _m.ctrl.Call(_m, "Apply", _param0)
	ret0, _ := ret[0].([]ApplyAttempt)
	ret1, _ := ret[1].([]ApplyAttempt)
	return ret0, ret1
}

func (_mr *_MockBatchApplierInterfaceRecorder) Apply(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Apply", arg0)
}
