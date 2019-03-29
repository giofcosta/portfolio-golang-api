package server

import (
	"github.com/giofcosta/portfolio-golang-api/config"
	"github.com/giofcosta/portfolio-golang-api/controllers"
	"github.com/giofcosta/portfolio-golang-api/repositories"
)

//Init initializes the server
func Init() {
	config := config.GetConfig()

	app := &Application{
		AboutMeController: &controllers.AboutMeController{
			Repository: repositories.NewAboutMeRepository(),
		},
		SkillController: &controllers.SkillController{
			Repository: repositories.NewSkillRepository(),
		},
		ResumeController: &controllers.ResumeController{
			Repository: repositories.NewResumeRepository(),
		},
	}

	r := NewRouter(app)

	r.Run(config.GetString("server.port"))
}
