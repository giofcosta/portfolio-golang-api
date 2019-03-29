package controllers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/giofcosta/portfolio-golang-api/mocks"

	"github.com/giofcosta/portfolio-golang-api/models"
	"github.com/giofcosta/portfolio-golang-api/server"
	"github.com/stretchr/testify/assert"
)

func TestAboutMeController_GET(t *testing.T) {
	appMock := mocks.NewApplicationMock(t)
	model := models.AboutMe{
		Items: []string{""},
	}

	appMock.MockAboutMeRepository.
		EXPECT().
		GetAll().
		Return(&model).
		AnyTimes()

	router := server.NewRouter(appMock.Mock)

	w := server.PerformRequest(router, "GET", "/v1/aboutme")

	var response models.AboutMe
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, model, response)
	assert.Nil(t, err)
	assert.NotNil(t, response.Items)
	assert.NotEmpty(t, response.Items)
}
