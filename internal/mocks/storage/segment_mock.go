// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/storage/segment.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"
	domain "segment-service/internal/core/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockSegmentStorager is a mock of SegmentStorager interface.
type MockSegmentStorager struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentStoragerMockRecorder
}

// MockSegmentStoragerMockRecorder is the mock recorder for MockSegmentStorager.
type MockSegmentStoragerMockRecorder struct {
	mock *MockSegmentStorager
}

// NewMockSegmentStorager creates a new mock instance.
func NewMockSegmentStorager(ctrl *gomock.Controller) *MockSegmentStorager {
	mock := &MockSegmentStorager{ctrl: ctrl}
	mock.recorder = &MockSegmentStoragerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSegmentStorager) EXPECT() *MockSegmentStoragerMockRecorder {
	return m.recorder
}

// AddSegment mocks base method.
func (m *MockSegmentStorager) AddSegment(arg0 context.Context, arg1 *domain.SegmentAddDTO) (*domain.Segment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSegment", arg0, arg1)
	ret0, _ := ret[0].(*domain.Segment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSegment indicates an expected call of AddSegment.
func (mr *MockSegmentStoragerMockRecorder) AddSegment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSegment", reflect.TypeOf((*MockSegmentStorager)(nil).AddSegment), arg0, arg1)
}

// CheckSegmentsExists mocks base method.
func (m *MockSegmentStorager) CheckSegmentsExists(arg0 context.Context, arg1 *domain.SegmentNames) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSegmentsExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckSegmentsExists indicates an expected call of CheckSegmentsExists.
func (mr *MockSegmentStoragerMockRecorder) CheckSegmentsExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSegmentsExists", reflect.TypeOf((*MockSegmentStorager)(nil).CheckSegmentsExists), arg0, arg1)
}

// DeleteSegment mocks base method.
func (m *MockSegmentStorager) DeleteSegment(arg0 context.Context, arg1 *domain.SegmentName) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSegment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSegment indicates an expected call of DeleteSegment.
func (mr *MockSegmentStoragerMockRecorder) DeleteSegment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSegment", reflect.TypeOf((*MockSegmentStorager)(nil).DeleteSegment), arg0, arg1)
}

// GetAllSegments mocks base method.
func (m *MockSegmentStorager) GetAllSegments(arg0 context.Context) (*[]domain.Segment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSegments", arg0)
	ret0, _ := ret[0].(*[]domain.Segment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSegments indicates an expected call of GetAllSegments.
func (mr *MockSegmentStoragerMockRecorder) GetAllSegments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSegments", reflect.TypeOf((*MockSegmentStorager)(nil).GetAllSegments), arg0)
}

// GetSegmentByName mocks base method.
func (m *MockSegmentStorager) GetSegmentByName(arg0 context.Context, arg1 *domain.SegmentName) (*domain.Segment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSegmentByName", arg0, arg1)
	ret0, _ := ret[0].(*domain.Segment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSegmentByName indicates an expected call of GetSegmentByName.
func (mr *MockSegmentStoragerMockRecorder) GetSegmentByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSegmentByName", reflect.TypeOf((*MockSegmentStorager)(nil).GetSegmentByName), arg0, arg1)
}

// GetSegmentsIds mocks base method.
func (m *MockSegmentStorager) GetSegmentsIds(ctx context.Context, dto *domain.SegmentNames) (*domain.SegmentIds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSegmentsIds", ctx, dto)
	ret0, _ := ret[0].(*domain.SegmentIds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSegmentsIds indicates an expected call of GetSegmentsIds.
func (mr *MockSegmentStoragerMockRecorder) GetSegmentsIds(ctx, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSegmentsIds", reflect.TypeOf((*MockSegmentStorager)(nil).GetSegmentsIds), ctx, dto)
}

// UpdateSegment mocks base method.
func (m *MockSegmentStorager) UpdateSegment(arg0 context.Context, arg1 *domain.SegmentUpdateDTO) (*domain.Segment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSegment", arg0, arg1)
	ret0, _ := ret[0].(*domain.Segment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSegment indicates an expected call of UpdateSegment.
func (mr *MockSegmentStoragerMockRecorder) UpdateSegment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSegment", reflect.TypeOf((*MockSegmentStorager)(nil).UpdateSegment), arg0, arg1)
}
