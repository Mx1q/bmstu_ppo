// Code generated by MockGen. DO NOT EDIT.
// Source: domain/measurement.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "ppo/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockIMeasurementRepository is a mock of IMeasurementRepository interface.
type MockIMeasurementRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIMeasurementRepositoryMockRecorder
}

// MockIMeasurementRepositoryMockRecorder is the mock recorder for MockIMeasurementRepository.
type MockIMeasurementRepositoryMockRecorder struct {
	mock *MockIMeasurementRepository
}

// NewMockIMeasurementRepository creates a new mock instance.
func NewMockIMeasurementRepository(ctrl *gomock.Controller) *MockIMeasurementRepository {
	mock := &MockIMeasurementRepository{ctrl: ctrl}
	mock.recorder = &MockIMeasurementRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMeasurementRepository) EXPECT() *MockIMeasurementRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIMeasurementRepository) Create(ctx context.Context, measurement *domain.Measurement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, measurement)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIMeasurementRepositoryMockRecorder) Create(ctx, measurement interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIMeasurementRepository)(nil).Create), ctx, measurement)
}

// DeleteById mocks base method.
func (m *MockIMeasurementRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockIMeasurementRepositoryMockRecorder) DeleteById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockIMeasurementRepository)(nil).DeleteById), ctx, id)
}

// GetAll mocks base method.
func (m *MockIMeasurementRepository) GetAll(ctx context.Context) ([]*domain.Measurement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*domain.Measurement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIMeasurementRepositoryMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIMeasurementRepository)(nil).GetAll), ctx)
}

// GetById mocks base method.
func (m *MockIMeasurementRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Measurement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*domain.Measurement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIMeasurementRepositoryMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIMeasurementRepository)(nil).GetById), ctx, id)
}

// GetByRecipeId mocks base method.
func (m *MockIMeasurementRepository) GetByRecipeId(ctx context.Context, ingredientId, recipeId uuid.UUID) (*domain.Measurement, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByRecipeId", ctx, ingredientId, recipeId)
	ret0, _ := ret[0].(*domain.Measurement)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByRecipeId indicates an expected call of GetByRecipeId.
func (mr *MockIMeasurementRepositoryMockRecorder) GetByRecipeId(ctx, ingredientId, recipeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByRecipeId", reflect.TypeOf((*MockIMeasurementRepository)(nil).GetByRecipeId), ctx, ingredientId, recipeId)
}

// Update mocks base method.
func (m *MockIMeasurementRepository) Update(ctx context.Context, measurement *domain.Measurement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, measurement)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIMeasurementRepositoryMockRecorder) Update(ctx, measurement interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIMeasurementRepository)(nil).Update), ctx, measurement)
}

// UpdateLink mocks base method.
func (m *MockIMeasurementRepository) UpdateLink(ctx context.Context, linkId, measurementId uuid.UUID, amount int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLink", ctx, linkId, measurementId, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLink indicates an expected call of UpdateLink.
func (mr *MockIMeasurementRepositoryMockRecorder) UpdateLink(ctx, linkId, measurementId, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLink", reflect.TypeOf((*MockIMeasurementRepository)(nil).UpdateLink), ctx, linkId, measurementId, amount)
}

// MockIMeasurementService is a mock of IMeasurementService interface.
type MockIMeasurementService struct {
	ctrl     *gomock.Controller
	recorder *MockIMeasurementServiceMockRecorder
}

// MockIMeasurementServiceMockRecorder is the mock recorder for MockIMeasurementService.
type MockIMeasurementServiceMockRecorder struct {
	mock *MockIMeasurementService
}

// NewMockIMeasurementService creates a new mock instance.
func NewMockIMeasurementService(ctrl *gomock.Controller) *MockIMeasurementService {
	mock := &MockIMeasurementService{ctrl: ctrl}
	mock.recorder = &MockIMeasurementServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMeasurementService) EXPECT() *MockIMeasurementServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIMeasurementService) Create(ctx context.Context, measurement *domain.Measurement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, measurement)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIMeasurementServiceMockRecorder) Create(ctx, measurement interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIMeasurementService)(nil).Create), ctx, measurement)
}

// DeleteById mocks base method.
func (m *MockIMeasurementService) DeleteById(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockIMeasurementServiceMockRecorder) DeleteById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockIMeasurementService)(nil).DeleteById), ctx, id)
}

// GetAll mocks base method.
func (m *MockIMeasurementService) GetAll(ctx context.Context) ([]*domain.Measurement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*domain.Measurement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIMeasurementServiceMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIMeasurementService)(nil).GetAll), ctx)
}

// GetById mocks base method.
func (m *MockIMeasurementService) GetById(ctx context.Context, id uuid.UUID) (*domain.Measurement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*domain.Measurement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIMeasurementServiceMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIMeasurementService)(nil).GetById), ctx, id)
}

// GetByRecipeId mocks base method.
func (m *MockIMeasurementService) GetByRecipeId(ctx context.Context, ingredientId, recipeId uuid.UUID) (*domain.Measurement, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByRecipeId", ctx, ingredientId, recipeId)
	ret0, _ := ret[0].(*domain.Measurement)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByRecipeId indicates an expected call of GetByRecipeId.
func (mr *MockIMeasurementServiceMockRecorder) GetByRecipeId(ctx, ingredientId, recipeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByRecipeId", reflect.TypeOf((*MockIMeasurementService)(nil).GetByRecipeId), ctx, ingredientId, recipeId)
}

// Update mocks base method.
func (m *MockIMeasurementService) Update(ctx context.Context, measurement *domain.Measurement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, measurement)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIMeasurementServiceMockRecorder) Update(ctx, measurement interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIMeasurementService)(nil).Update), ctx, measurement)
}

// UpdateLink mocks base method.
func (m *MockIMeasurementService) UpdateLink(ctx context.Context, linkId, measurementId uuid.UUID, amount int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLink", ctx, linkId, measurementId, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLink indicates an expected call of UpdateLink.
func (mr *MockIMeasurementServiceMockRecorder) UpdateLink(ctx, linkId, measurementId, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLink", reflect.TypeOf((*MockIMeasurementService)(nil).UpdateLink), ctx, linkId, measurementId, amount)
}
