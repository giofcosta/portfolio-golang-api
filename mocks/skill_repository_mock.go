package mocks

import (
	reflect "reflect"

	models "github.com/giofcosta/portfolio-golang-api/models"
	gomock "github.com/golang/mock/gomock"
)

// MockSkillRepository is a mock of SkillRepository interface
type MockSkillRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSkillRepositoryMockRecorder
}

// MockSkillRepositoryMockRecorder is the mock recorder for MockSkillRepository
type MockSkillRepositoryMockRecorder struct {
	mock *MockSkillRepository
}

// NewMockSkillRepository creates a new mock instance
func NewMockSkillRepository(ctrl *gomock.Controller) *MockSkillRepository {
	mock := &MockSkillRepository{ctrl: ctrl}
	mock.recorder = &MockSkillRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSkillRepository) EXPECT() *MockSkillRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockSkillRepository) GetAll() *models.Skill {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*models.Skill)
	return ret0
}

// GetAll indicates an expected call of GetAll
func (mr *MockSkillRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockSkillRepository)(nil).GetAll))
}

// GetItems mocks base method
func (m *MockSkillRepository) GetItems() []models.Items {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems")
	ret0, _ := ret[0].([]models.Items)
	return ret0
}

// GetItems indicates an expected call of GetItems
func (mr *MockSkillRepositoryMockRecorder) GetItems() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockSkillRepository)(nil).GetItems))
}

// GetLanguages mocks base method
func (m *MockSkillRepository) GetLanguages() []models.Languages {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLanguages")
	ret0, _ := ret[0].([]models.Languages)
	return ret0
}

// GetLanguages indicates an expected call of GetLanguages
func (mr *MockSkillRepositoryMockRecorder) GetLanguages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLanguages", reflect.TypeOf((*MockSkillRepository)(nil).GetLanguages))
}

// GetTags mocks base method
func (m *MockSkillRepository) GetTags() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTags")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetTags indicates an expected call of GetTags
func (mr *MockSkillRepositoryMockRecorder) GetTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockSkillRepository)(nil).GetTags))
}
