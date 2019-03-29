package mocks

import (
	"testing"

	"github.com/giofcosta/portfolio-golang-api/controllers"
	"github.com/giofcosta/portfolio-golang-api/server"
	gomock "github.com/golang/mock/gomock"
)

type ApplicationMock struct {
	Mock                  *server.Application
	Ctrl                  *gomock.Controller
	MockAboutMeRepository *MockAboutMeRepository
	MockSkillRepository   *MockSkillRepository
	MockResumeRepository  *MockResumeRepository
}

func NewApplicationMock(t *testing.T) *ApplicationMock {
	ctrl := gomock.NewController(t)

	appMock := &ApplicationMock{
		Ctrl: ctrl,
		MockAboutMeRepository: NewMockAboutMeRepository(ctrl),
		MockResumeRepository:  NewMockResumeRepository(ctrl),
		MockSkillRepository:   NewMockSkillRepository(ctrl),
	}

	defer appMock.Ctrl.Finish()

	app := server.BuildApplication()
	app.AboutMeController.Repository = appMock.MockAboutMeRepository
	app.SkillController.Repository = appMock.MockSkillRepository
	app.ResumeController.Repository = appMock.MockResumeRepository

	appMock.Mock = &server.Application{
		AboutMeController: &controllers.AboutMeController{
			Repository: appMock.MockAboutMeRepository,
		},
		SkillController: &controllers.SkillController{
			Repository: appMock.MockSkillRepository,
		},
		ResumeController: &controllers.ResumeController{
			Repository: appMock.MockResumeRepository,
		},
	}

	return appMock
}
