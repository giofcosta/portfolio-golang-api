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

type RepositoryMock struct {
}

func TestAboutMeController_GET(t *testing.T) {
	//Mocked app
	app := &server.Application{
		AboutMeController: &controllers.AboutMeController{
			//Repository: RepositoryMock{},
		},
	}

	router := server.NewRouter(app)

	w := utils.PerformRequest(router, "GET", "/aboutme")

	var response models.AboutMe
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response.Items)
}
