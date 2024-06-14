// Code generated by MockGen. DO NOT EDIT.
// Source: domain/saladInteractor.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "ppo/domain"
	dto "ppo/services/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockISaladInteractor is a mock of ISaladInteractor interface.
type MockISaladInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockISaladInteractorMockRecorder
}

// MockISaladInteractorMockRecorder is the mock recorder for MockISaladInteractor.
type MockISaladInteractorMockRecorder struct {
	mock *MockISaladInteractor
}

// NewMockISaladInteractor creates a new mock instance.
func NewMockISaladInteractor(ctrl *gomock.Controller) *MockISaladInteractor {
	mock := &MockISaladInteractor{ctrl: ctrl}
	mock.recorder = &MockISaladInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISaladInteractor) EXPECT() *MockISaladInteractorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockISaladInteractor) Create(ctx context.Context, salad *domain.Salad) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, salad)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockISaladInteractorMockRecorder) Create(ctx, salad interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockISaladInteractor)(nil).Create), ctx, salad)
}

// DeleteById mocks base method.
func (m *MockISaladInteractor) DeleteById(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockISaladInteractorMockRecorder) DeleteById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockISaladInteractor)(nil).DeleteById), ctx, id)
}

// GetAll mocks base method.
func (m *MockISaladInteractor) GetAll(ctx context.Context, filter *dto.RecipeFilter, page int) ([]*domain.Salad, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, filter, page)
	ret0, _ := ret[0].([]*domain.Salad)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAll indicates an expected call of GetAll.
func (mr *MockISaladInteractorMockRecorder) GetAll(ctx, filter, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockISaladInteractor)(nil).GetAll), ctx, filter, page)
}

// GetAllByUserId mocks base method.
func (m *MockISaladInteractor) GetAllByUserId(ctx context.Context, id uuid.UUID) ([]*domain.Salad, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByUserId", ctx, id)
	ret0, _ := ret[0].([]*domain.Salad)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByUserId indicates an expected call of GetAllByUserId.
func (mr *MockISaladInteractorMockRecorder) GetAllByUserId(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByUserId", reflect.TypeOf((*MockISaladInteractor)(nil).GetAllByUserId), ctx, id)
}

// GetAllRatedByUser mocks base method.
func (m *MockISaladInteractor) GetAllRatedByUser(ctx context.Context, userId uuid.UUID, page int) ([]*domain.Salad, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRatedByUser", ctx, userId, page)
	ret0, _ := ret[0].([]*domain.Salad)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllRatedByUser indicates an expected call of GetAllRatedByUser.
func (mr *MockISaladInteractorMockRecorder) GetAllRatedByUser(ctx, userId, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRatedByUser", reflect.TypeOf((*MockISaladInteractor)(nil).GetAllRatedByUser), ctx, userId, page)
}

// GetById mocks base method.
func (m *MockISaladInteractor) GetById(ctx context.Context, id uuid.UUID) (*domain.Salad, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*domain.Salad)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockISaladInteractorMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockISaladInteractor)(nil).GetById), ctx, id)
}

// Update mocks base method.
func (m *MockISaladInteractor) Update(ctx context.Context, salad *domain.Salad) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, salad)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockISaladInteractorMockRecorder) Update(ctx, salad interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockISaladInteractor)(nil).Update), ctx, salad)
}
