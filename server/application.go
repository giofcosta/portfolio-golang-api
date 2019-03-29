package server

import "github.com/giofcosta/portfolio-golang-api/controllers"

type Application struct {
	AboutMeController *controllers.AboutMeController
	SkillController   *controllers.SkillController
	ResumeController  *controllers.ResumeController
}
