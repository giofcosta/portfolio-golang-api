package server

import (
	"net/http"
	"net/http/httptest"

	"github.com/giofcosta/portfolio-golang-api/config"
	"github.com/giofcosta/portfolio-golang-api/controllers"
	"github.com/giofcosta/portfolio-golang-api/repositories"
)

//Init initializes the server
func Init() {
	config := config.GetConfig()

	r := NewRouter(BuildApplication())

	r.Run(config.GetString("server.port"))
}

//BuildApplication builds the application instance
func BuildApplication() *Application {
	return &Application{
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
}

func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
