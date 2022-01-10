/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by MockGen. DO NOT EDIT.
// Source: ./powervs.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/IBM-Cloud/power-go-client/power/models"
	gomock "github.com/golang/mock/gomock"
)

// MockPowerVS is a mock of PowerVS interface.
type MockPowerVS struct {
	ctrl     *gomock.Controller
	recorder *MockPowerVSMockRecorder
}

// MockPowerVSMockRecorder is the mock recorder for MockPowerVS.
type MockPowerVSMockRecorder struct {
	mock *MockPowerVS
}

// NewMockPowerVS creates a new mock instance.
func NewMockPowerVS(ctrl *gomock.Controller) *MockPowerVS {
	mock := &MockPowerVS{ctrl: ctrl}
	mock.recorder = &MockPowerVSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPowerVS) EXPECT() *MockPowerVSMockRecorder {
	return m.recorder
}

// CreateInstance mocks base method.
func (m *MockPowerVS) CreateInstance(body *models.PVMInstanceCreate) (*models.PVMInstanceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstance", body)
	ret0, _ := ret[0].(*models.PVMInstanceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInstance indicates an expected call of CreateInstance.
func (mr *MockPowerVSMockRecorder) CreateInstance(body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstance", reflect.TypeOf((*MockPowerVS)(nil).CreateInstance), body)
}

// DeleteInstance mocks base method.
func (m *MockPowerVS) DeleteInstance(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInstance", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInstance indicates an expected call of DeleteInstance.
func (mr *MockPowerVSMockRecorder) DeleteInstance(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstance", reflect.TypeOf((*MockPowerVS)(nil).DeleteInstance), id)
}

// GetAllImage mocks base method.
func (m *MockPowerVS) GetAllImage() (*models.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllImage")
	ret0, _ := ret[0].(*models.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllImage indicates an expected call of GetAllImage.
func (mr *MockPowerVSMockRecorder) GetAllImage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllImage", reflect.TypeOf((*MockPowerVS)(nil).GetAllImage))
}

// GetAllInstance mocks base method.
func (m *MockPowerVS) GetAllInstance() (*models.PVMInstances, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllInstance")
	ret0, _ := ret[0].(*models.PVMInstances)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllInstance indicates an expected call of GetAllInstance.
func (mr *MockPowerVSMockRecorder) GetAllInstance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllInstance", reflect.TypeOf((*MockPowerVS)(nil).GetAllInstance))
}

// GetAllNetwork mocks base method.
func (m *MockPowerVS) GetAllNetwork() (*models.Networks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllNetwork")
	ret0, _ := ret[0].(*models.Networks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllNetwork indicates an expected call of GetAllNetwork.
func (mr *MockPowerVSMockRecorder) GetAllNetwork() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllNetwork", reflect.TypeOf((*MockPowerVS)(nil).GetAllNetwork))
}
