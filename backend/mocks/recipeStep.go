// Code generated by MockGen. DO NOT EDIT.
// Source: domain/recipeStep.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "ppo/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockIRecipeStepRepository is a mock of IRecipeStepRepository interface.
type MockIRecipeStepRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRecipeStepRepositoryMockRecorder
}

// MockIRecipeStepRepositoryMockRecorder is the mock recorder for MockIRecipeStepRepository.
type MockIRecipeStepRepositoryMockRecorder struct {
	mock *MockIRecipeStepRepository
}

// NewMockIRecipeStepRepository creates a new mock instance.
func NewMockIRecipeStepRepository(ctrl *gomock.Controller) *MockIRecipeStepRepository {
	mock := &MockIRecipeStepRepository{ctrl: ctrl}
	mock.recorder = &MockIRecipeStepRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRecipeStepRepository) EXPECT() *MockIRecipeStepRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIRecipeStepRepository) Create(ctx context.Context, ingredient *domain.RecipeStep) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, ingredient)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIRecipeStepRepositoryMockRecorder) Create(ctx, ingredient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIRecipeStepRepository)(nil).Create), ctx, ingredient)
}

// DeleteAllByRecipeID mocks base method.
func (m *MockIRecipeStepRepository) DeleteAllByRecipeID(ctx context.Context, recipeId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllByRecipeID", ctx, recipeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllByRecipeID indicates an expected call of DeleteAllByRecipeID.
func (mr *MockIRecipeStepRepositoryMockRecorder) DeleteAllByRecipeID(ctx, recipeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllByRecipeID", reflect.TypeOf((*MockIRecipeStepRepository)(nil).DeleteAllByRecipeID), ctx, recipeId)
}

// DeleteById mocks base method.
func (m *MockIRecipeStepRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockIRecipeStepRepositoryMockRecorder) DeleteById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockIRecipeStepRepository)(nil).DeleteById), ctx, id)
}

// GetAllByRecipeID mocks base method.
func (m *MockIRecipeStepRepository) GetAllByRecipeID(ctx context.Context, recipeId uuid.UUID) ([]*domain.RecipeStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByRecipeID", ctx, recipeId)
	ret0, _ := ret[0].([]*domain.RecipeStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByRecipeID indicates an expected call of GetAllByRecipeID.
func (mr *MockIRecipeStepRepositoryMockRecorder) GetAllByRecipeID(ctx, recipeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByRecipeID", reflect.TypeOf((*MockIRecipeStepRepository)(nil).GetAllByRecipeID), ctx, recipeId)
}

// GetById mocks base method.
func (m *MockIRecipeStepRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.RecipeStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*domain.RecipeStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIRecipeStepRepositoryMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIRecipeStepRepository)(nil).GetById), ctx, id)
}

// Update mocks base method.
func (m *MockIRecipeStepRepository) Update(ctx context.Context, ingredient *domain.RecipeStep) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, ingredient)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIRecipeStepRepositoryMockRecorder) Update(ctx, ingredient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRecipeStepRepository)(nil).Update), ctx, ingredient)
}

// MockIRecipeStepService is a mock of IRecipeStepService interface.
type MockIRecipeStepService struct {
	ctrl     *gomock.Controller
	recorder *MockIRecipeStepServiceMockRecorder
}

// MockIRecipeStepServiceMockRecorder is the mock recorder for MockIRecipeStepService.
type MockIRecipeStepServiceMockRecorder struct {
	mock *MockIRecipeStepService
}

// NewMockIRecipeStepService creates a new mock instance.
func NewMockIRecipeStepService(ctrl *gomock.Controller) *MockIRecipeStepService {
	mock := &MockIRecipeStepService{ctrl: ctrl}
	mock.recorder = &MockIRecipeStepServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRecipeStepService) EXPECT() *MockIRecipeStepServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIRecipeStepService) Create(ctx context.Context, recipeStep *domain.RecipeStep) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, recipeStep)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIRecipeStepServiceMockRecorder) Create(ctx, recipeStep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIRecipeStepService)(nil).Create), ctx, recipeStep)
}

// DeleteAllByRecipeID mocks base method.
func (m *MockIRecipeStepService) DeleteAllByRecipeID(ctx context.Context, recipeId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllByRecipeID", ctx, recipeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllByRecipeID indicates an expected call of DeleteAllByRecipeID.
func (mr *MockIRecipeStepServiceMockRecorder) DeleteAllByRecipeID(ctx, recipeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllByRecipeID", reflect.TypeOf((*MockIRecipeStepService)(nil).DeleteAllByRecipeID), ctx, recipeId)
}

// DeleteById mocks base method.
func (m *MockIRecipeStepService) DeleteById(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockIRecipeStepServiceMockRecorder) DeleteById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockIRecipeStepService)(nil).DeleteById), ctx, id)
}

// GetAllByRecipeID mocks base method.
func (m *MockIRecipeStepService) GetAllByRecipeID(ctx context.Context, recipeId uuid.UUID) ([]*domain.RecipeStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByRecipeID", ctx, recipeId)
	ret0, _ := ret[0].([]*domain.RecipeStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByRecipeID indicates an expected call of GetAllByRecipeID.
func (mr *MockIRecipeStepServiceMockRecorder) GetAllByRecipeID(ctx, recipeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByRecipeID", reflect.TypeOf((*MockIRecipeStepService)(nil).GetAllByRecipeID), ctx, recipeId)
}

// GetById mocks base method.
func (m *MockIRecipeStepService) GetById(ctx context.Context, id uuid.UUID) (*domain.RecipeStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*domain.RecipeStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIRecipeStepServiceMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIRecipeStepService)(nil).GetById), ctx, id)
}

// Update mocks base method.
func (m *MockIRecipeStepService) Update(ctx context.Context, recipeStep *domain.RecipeStep) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, recipeStep)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIRecipeStepServiceMockRecorder) Update(ctx, recipeStep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRecipeStepService)(nil).Update), ctx, recipeStep)
}
