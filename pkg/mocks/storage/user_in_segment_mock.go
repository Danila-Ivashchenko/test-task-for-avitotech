// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/storage/user-in-segment.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"
	domain "segment-service/internal/core/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockUserInSegmentStorager is a mock of UserInSegmentStorager interface.
type MockUserInSegmentStorager struct {
	ctrl     *gomock.Controller
	recorder *MockUserInSegmentStoragerMockRecorder
}

// MockUserInSegmentStoragerMockRecorder is the mock recorder for MockUserInSegmentStorager.
type MockUserInSegmentStoragerMockRecorder struct {
	mock *MockUserInSegmentStorager
}

// NewMockUserInSegmentStorager creates a new mock instance.
func NewMockUserInSegmentStorager(ctrl *gomock.Controller) *MockUserInSegmentStorager {
	mock := &MockUserInSegmentStorager{ctrl: ctrl}
	mock.recorder = &MockUserInSegmentStoragerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserInSegmentStorager) EXPECT() *MockUserInSegmentStoragerMockRecorder {
	return m.recorder
}

// AddPercentOfUsersToSegments mocks base method.
func (m *MockUserInSegmentStorager) AddPercentOfUsersToSegments(arg0 context.Context, arg1 *domain.PercentOfUsersToSegmentsDTO) (*domain.UsersIds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPercentOfUsersToSegments", arg0, arg1)
	ret0, _ := ret[0].(*domain.UsersIds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPercentOfUsersToSegments indicates an expected call of AddPercentOfUsersToSegments.
func (mr *MockUserInSegmentStoragerMockRecorder) AddPercentOfUsersToSegments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPercentOfUsersToSegments", reflect.TypeOf((*MockUserInSegmentStorager)(nil).AddPercentOfUsersToSegments), arg0, arg1)
}

// AddUserToSegments mocks base method.
func (m *MockUserInSegmentStorager) AddUserToSegments(arg0 context.Context, arg1 *domain.UserToSegmentsAddDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserToSegments", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserToSegments indicates an expected call of AddUserToSegments.
func (mr *MockUserInSegmentStoragerMockRecorder) AddUserToSegments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToSegments", reflect.TypeOf((*MockUserInSegmentStorager)(nil).AddUserToSegments), arg0, arg1)
}

// AddUsersToSegments mocks base method.
func (m *MockUserInSegmentStorager) AddUsersToSegments(arg0 context.Context, arg1 *domain.UsersToSegmentsAddDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUsersToSegments", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUsersToSegments indicates an expected call of AddUsersToSegments.
func (mr *MockUserInSegmentStoragerMockRecorder) AddUsersToSegments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUsersToSegments", reflect.TypeOf((*MockUserInSegmentStorager)(nil).AddUsersToSegments), arg0, arg1)
}

// AddUsersWithLimitOffsetToSegments mocks base method.
func (m *MockUserInSegmentStorager) AddUsersWithLimitOffsetToSegments(arg0 context.Context, arg1 *domain.UsersWithLimitOffsetToSegments) (*domain.UsersIds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUsersWithLimitOffsetToSegments", arg0, arg1)
	ret0, _ := ret[0].(*domain.UsersIds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUsersWithLimitOffsetToSegments indicates an expected call of AddUsersWithLimitOffsetToSegments.
func (mr *MockUserInSegmentStoragerMockRecorder) AddUsersWithLimitOffsetToSegments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUsersWithLimitOffsetToSegments", reflect.TypeOf((*MockUserInSegmentStorager)(nil).AddUsersWithLimitOffsetToSegments), arg0, arg1)
}

// DeleteUserFromSegments mocks base method.
func (m *MockUserInSegmentStorager) DeleteUserFromSegments(arg0 context.Context, arg1 *domain.UserFromSegmentsDeleteDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserFromSegments", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserFromSegments indicates an expected call of DeleteUserFromSegments.
func (mr *MockUserInSegmentStoragerMockRecorder) DeleteUserFromSegments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserFromSegments", reflect.TypeOf((*MockUserInSegmentStorager)(nil).DeleteUserFromSegments), arg0, arg1)
}

// GetUserInSegments mocks base method.
func (m *MockUserInSegmentStorager) GetUserInSegments(arg0 context.Context, arg1 *domain.UserId) (*domain.UserInSegments, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInSegments", arg0, arg1)
	ret0, _ := ret[0].(*domain.UserInSegments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInSegments indicates an expected call of GetUserInSegments.
func (mr *MockUserInSegmentStoragerMockRecorder) GetUserInSegments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInSegments", reflect.TypeOf((*MockUserInSegmentStorager)(nil).GetUserInSegments), arg0, arg1)
}

// GetUsersInSegment mocks base method.
func (m *MockUserInSegmentStorager) GetUsersInSegment(arg0 context.Context, arg1 *domain.SegmentName) (*domain.UsersInSegment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersInSegment", arg0, arg1)
	ret0, _ := ret[0].(*domain.UsersInSegment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersInSegment indicates an expected call of GetUsersInSegment.
func (mr *MockUserInSegmentStoragerMockRecorder) GetUsersInSegment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersInSegment", reflect.TypeOf((*MockUserInSegmentStorager)(nil).GetUsersInSegment), arg0, arg1)
}
