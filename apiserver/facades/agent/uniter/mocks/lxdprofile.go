// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/uniter (interfaces: LXDProfileBackend,LXDProfileMachine,LXDProfileUnit)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	uniter "github.com/juju/juju/apiserver/facades/agent/uniter"
	state "github.com/juju/juju/state"
	names_v2 "gopkg.in/juju/names.v2"
	reflect "reflect"
)

// MockLXDProfileBackend is a mock of LXDProfileBackend interface
type MockLXDProfileBackend struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileBackendMockRecorder
}

// MockLXDProfileBackendMockRecorder is the mock recorder for MockLXDProfileBackend
type MockLXDProfileBackendMockRecorder struct {
	mock *MockLXDProfileBackend
}

// NewMockLXDProfileBackend creates a new mock instance
func NewMockLXDProfileBackend(ctrl *gomock.Controller) *MockLXDProfileBackend {
	mock := &MockLXDProfileBackend{ctrl: ctrl}
	mock.recorder = &MockLXDProfileBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLXDProfileBackend) EXPECT() *MockLXDProfileBackendMockRecorder {
	return m.recorder
}

// Machine mocks base method
func (m *MockLXDProfileBackend) Machine(arg0 string) (uniter.LXDProfileMachine, error) {
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(uniter.LXDProfileMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine
func (mr *MockLXDProfileBackendMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockLXDProfileBackend)(nil).Machine), arg0)
}

// Unit mocks base method
func (m *MockLXDProfileBackend) Unit(arg0 string) (uniter.LXDProfileUnit, error) {
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(uniter.LXDProfileUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit
func (mr *MockLXDProfileBackendMockRecorder) Unit(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockLXDProfileBackend)(nil).Unit), arg0)
}

// MockLXDProfileMachine is a mock of LXDProfileMachine interface
type MockLXDProfileMachine struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileMachineMockRecorder
}

// MockLXDProfileMachineMockRecorder is the mock recorder for MockLXDProfileMachine
type MockLXDProfileMachineMockRecorder struct {
	mock *MockLXDProfileMachine
}

// NewMockLXDProfileMachine creates a new mock instance
func NewMockLXDProfileMachine(ctrl *gomock.Controller) *MockLXDProfileMachine {
	mock := &MockLXDProfileMachine{ctrl: ctrl}
	mock.recorder = &MockLXDProfileMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLXDProfileMachine) EXPECT() *MockLXDProfileMachineMockRecorder {
	return m.recorder
}

// RemoveUpgradeCharmProfileData mocks base method
func (m *MockLXDProfileMachine) RemoveUpgradeCharmProfileData(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveUpgradeCharmProfileData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUpgradeCharmProfileData indicates an expected call of RemoveUpgradeCharmProfileData
func (mr *MockLXDProfileMachineMockRecorder) RemoveUpgradeCharmProfileData(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUpgradeCharmProfileData", reflect.TypeOf((*MockLXDProfileMachine)(nil).RemoveUpgradeCharmProfileData), arg0)
}

// Units mocks base method
func (m *MockLXDProfileMachine) Units() ([]uniter.LXDProfileUnit, error) {
	ret := m.ctrl.Call(m, "Units")
	ret0, _ := ret[0].([]uniter.LXDProfileUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units
func (mr *MockLXDProfileMachineMockRecorder) Units() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockLXDProfileMachine)(nil).Units))
}

// WatchLXDProfileUpgradeNotifications mocks base method
func (m *MockLXDProfileMachine) WatchLXDProfileUpgradeNotifications(arg0 string) (state.StringsWatcher, error) {
	ret := m.ctrl.Call(m, "WatchLXDProfileUpgradeNotifications", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchLXDProfileUpgradeNotifications indicates an expected call of WatchLXDProfileUpgradeNotifications
func (mr *MockLXDProfileMachineMockRecorder) WatchLXDProfileUpgradeNotifications(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchLXDProfileUpgradeNotifications", reflect.TypeOf((*MockLXDProfileMachine)(nil).WatchLXDProfileUpgradeNotifications), arg0)
}

// MockLXDProfileUnit is a mock of LXDProfileUnit interface
type MockLXDProfileUnit struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfileUnitMockRecorder
}

// MockLXDProfileUnitMockRecorder is the mock recorder for MockLXDProfileUnit
type MockLXDProfileUnitMockRecorder struct {
	mock *MockLXDProfileUnit
}

// NewMockLXDProfileUnit creates a new mock instance
func NewMockLXDProfileUnit(ctrl *gomock.Controller) *MockLXDProfileUnit {
	mock := &MockLXDProfileUnit{ctrl: ctrl}
	mock.recorder = &MockLXDProfileUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLXDProfileUnit) EXPECT() *MockLXDProfileUnitMockRecorder {
	return m.recorder
}

// AssignedMachineId mocks base method
func (m *MockLXDProfileUnit) AssignedMachineId() (string, error) {
	ret := m.ctrl.Call(m, "AssignedMachineId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignedMachineId indicates an expected call of AssignedMachineId
func (mr *MockLXDProfileUnitMockRecorder) AssignedMachineId() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignedMachineId", reflect.TypeOf((*MockLXDProfileUnit)(nil).AssignedMachineId))
}

// Name mocks base method
func (m *MockLXDProfileUnit) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockLXDProfileUnitMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockLXDProfileUnit)(nil).Name))
}

// Tag mocks base method
func (m *MockLXDProfileUnit) Tag() names_v2.Tag {
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names_v2.Tag)
	return ret0
}

// Tag indicates an expected call of Tag
func (mr *MockLXDProfileUnitMockRecorder) Tag() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockLXDProfileUnit)(nil).Tag))
}
