package mocks

import (
	reflect "reflect"

	models "github.com/giofcosta/portfolio-golang-api/models"
	gomock "github.com/golang/mock/gomock"
)

// MockAboutMeRepository is a mock of AboutMeRepository interface
type MockAboutMeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAboutMeRepositoryMockRecorder
}

// MockAboutMeRepositoryMockRecorder is the mock recorder for MockAboutMeRepository
type MockAboutMeRepositoryMockRecorder struct {
	mock *MockAboutMeRepository
}

// NewMockAboutMeRepository creates a new mock instance
func NewMockAboutMeRepository(ctrl *gomock.Controller) *MockAboutMeRepository {
	mock := &MockAboutMeRepository{ctrl: ctrl}
	mock.recorder = &MockAboutMeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAboutMeRepository) EXPECT() *MockAboutMeRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockAboutMeRepository) GetAll() *models.AboutMe {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*models.AboutMe)
	return ret0
}

// GetAll indicates an expected call of GetAll
func (mr *MockAboutMeRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAboutMeRepository)(nil).GetAll))
}

// GetItems mocks base method
func (m *MockAboutMeRepository) GetItems() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetItems indicates an expected call of GetItems
func (mr *MockAboutMeRepositoryMockRecorder) GetItems() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockAboutMeRepository)(nil).GetItems))
}
