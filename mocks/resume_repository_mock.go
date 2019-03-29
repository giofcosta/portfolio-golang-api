package mocks

import (
	reflect "reflect"

	models "github.com/giofcosta/portfolio-golang-api/models"
	gomock "github.com/golang/mock/gomock"
)

// MockResumeRepository is a mock of ResumeRepository interface
type MockResumeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockResumeRepositoryMockRecorder
}

// MockResumeRepositoryMockRecorder is the mock recorder for MockResumeRepository
type MockResumeRepositoryMockRecorder struct {
	mock *MockResumeRepository
}

// NewMockResumeRepository creates a new mock instance
func NewMockResumeRepository(ctrl *gomock.Controller) *MockResumeRepository {
	mock := &MockResumeRepository{ctrl: ctrl}
	mock.recorder = &MockResumeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResumeRepository) EXPECT() *MockResumeRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockResumeRepository) GetAll() *models.Resume {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*models.Resume)
	return ret0
}

// GetAll indicates an expected call of GetAll
func (mr *MockResumeRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockResumeRepository)(nil).GetAll))
}

// GetExperiences mocks base method
func (m *MockResumeRepository) GetExperiences() []models.Experiences {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExperiences")
	ret0, _ := ret[0].([]models.Experiences)
	return ret0
}

// GetExperiences indicates an expected call of GetExperiences
func (mr *MockResumeRepositoryMockRecorder) GetExperiences() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExperiences", reflect.TypeOf((*MockResumeRepository)(nil).GetExperiences))
}

// GetEducation mocks base method
func (m *MockResumeRepository) GetEducation() []models.Education {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEducation")
	ret0, _ := ret[0].([]models.Education)
	return ret0
}

// GetEducation indicates an expected call of GetEducation
func (mr *MockResumeRepositoryMockRecorder) GetEducation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEducation", reflect.TypeOf((*MockResumeRepository)(nil).GetEducation))
}

// GetCertifications mocks base method
func (m *MockResumeRepository) GetCertifications() []models.Certifications {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertifications")
	ret0, _ := ret[0].([]models.Certifications)
	return ret0
}

// GetCertifications indicates an expected call of GetCertifications
func (mr *MockResumeRepositoryMockRecorder) GetCertifications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertifications", reflect.TypeOf((*MockResumeRepository)(nil).GetCertifications))
}

// GetReadings mocks base method
func (m *MockResumeRepository) GetReadings() []models.Readings {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadings")
	ret0, _ := ret[0].([]models.Readings)
	return ret0
}

// GetReadings indicates an expected call of GetReadings
func (mr *MockResumeRepositoryMockRecorder) GetReadings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadings", reflect.TypeOf((*MockResumeRepository)(nil).GetReadings))
}
