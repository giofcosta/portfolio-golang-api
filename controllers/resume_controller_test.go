package controllers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/giofcosta/portfolio-golang-api/controllers"
	"github.com/giofcosta/portfolio-golang-api/utils"

	"github.com/giofcosta/portfolio-golang-api/models"
	"github.com/giofcosta/portfolio-golang-api/server"
	"github.com/stretchr/testify/assert"
)

func TestResumeController_GET(t *testing.T) {
	app := &server.Application{
		ResumeController: &controllers.ResumeController{
			//Repository: RepositoryMock{},
		},
	}

	router := server.NewRouter(app)

	w := utils.PerformRequest(router, "GET", "/resume")

	var response models.Resume
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response.Certifications)
	assert.NotNil(t, response.Education)
	assert.NotNil(t, response.Experiences)
	assert.NotNil(t, response.Readings)
}
