// Code generated by MockGen. DO NOT EDIT.
// Source: src/internal/core/ports/services.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"

	domain "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockSchedulerService is a mock of SchedulerService interface.
type MockSchedulerService struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerServiceMockRecorder
}

// MockSchedulerServiceMockRecorder is the mock recorder for MockSchedulerService.
type MockSchedulerServiceMockRecorder struct {
	mock *MockSchedulerService
}

// NewMockSchedulerService creates a new mock instance.
func NewMockSchedulerService(ctrl *gomock.Controller) *MockSchedulerService {
	mock := &MockSchedulerService{ctrl: ctrl}
	mock.recorder = &MockSchedulerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchedulerService) EXPECT() *MockSchedulerServiceMockRecorder {
	return m.recorder
}

// GetAvailableHours mocks base method.
func (m *MockSchedulerService) GetAvailableHours(terna domain.DegreeSet) ([]domain.AvailableHours, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableHours", terna)
	ret0, _ := ret[0].([]domain.AvailableHours)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableHours indicates an expected call of GetAvailableHours.
func (mr *MockSchedulerServiceMockRecorder) GetAvailableHours(terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableHours", reflect.TypeOf((*MockSchedulerService)(nil).GetAvailableHours), terna)
}

// GetEntries mocks base method.
func (m *MockSchedulerService) GetEntries(terna domain.DegreeSet) ([]domain.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntries", terna)
	ret0, _ := ret[0].([]domain.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntries indicates an expected call of GetEntries.
func (mr *MockSchedulerServiceMockRecorder) GetEntries(terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntries", reflect.TypeOf((*MockSchedulerService)(nil).GetEntries), terna)
}

// GetICS mocks base method.
func (m *MockSchedulerService) GetICS(terna domain.DegreeSet) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetICS", terna)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetICS indicates an expected call of GetICS.
func (mr *MockSchedulerServiceMockRecorder) GetICS(terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetICS", reflect.TypeOf((*MockSchedulerService)(nil).GetICS), terna)
}

// ListAllDegrees mocks base method.
func (m *MockSchedulerService) ListAllDegrees() ([]domain.DegreeDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllDegrees")
	ret0, _ := ret[0].([]domain.DegreeDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllDegrees indicates an expected call of ListAllDegrees.
func (mr *MockSchedulerServiceMockRecorder) ListAllDegrees() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllDegrees", reflect.TypeOf((*MockSchedulerService)(nil).ListAllDegrees))
}

// UpdateScheduler mocks base method.
func (m *MockSchedulerService) UpdateScheduler(entries []domain.Entry, terna domain.DegreeSet) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScheduler", entries, terna)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScheduler indicates an expected call of UpdateScheduler.
func (mr *MockSchedulerServiceMockRecorder) UpdateScheduler(entries, terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScheduler", reflect.TypeOf((*MockSchedulerService)(nil).UpdateScheduler), entries, terna)
}

// MockUploadDataservice is a mock of UploadDataservice interface.
type MockUploadDataservice struct {
	ctrl     *gomock.Controller
	recorder *MockUploadDataserviceMockRecorder
}

// MockUploadDataserviceMockRecorder is the mock recorder for MockUploadDataservice.
type MockUploadDataserviceMockRecorder struct {
	mock *MockUploadDataservice
}

// NewMockUploadDataservice creates a new mock instance.
func NewMockUploadDataservice(ctrl *gomock.Controller) *MockUploadDataservice {
	mock := &MockUploadDataservice{ctrl: ctrl}
	mock.recorder = &MockUploadDataserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadDataservice) EXPECT() *MockUploadDataserviceMockRecorder {
	return m.recorder
}

// UpdateByCSV mocks base method.
func (m *MockUploadDataservice) UpdateByCSV(csv string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByCSV", csv)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByCSV indicates an expected call of UpdateByCSV.
func (mr *MockUploadDataserviceMockRecorder) UpdateByCSV(csv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByCSV", reflect.TypeOf((*MockUploadDataservice)(nil).UpdateByCSV), csv)
}
