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

func TestResumeController_GET(t *testing.T) {
	appMock := mocks.NewApplicationMock(t)
	model := models.Resume{
		Certifications: []models.Certifications{},
		Education:      []models.Education{},
		Experiences:    []models.Experiences{},
		Readings:       []models.Readings{},
	}

	appMock.MockResumeRepository.
		EXPECT().
		GetAll().
		Return(&model).
		AnyTimes()

	router := server.NewRouter(appMock.Mock)

	w := server.PerformRequest(router, "GET", "/v1/resume")

	var response models.Resume
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, model, response)
	assert.Nil(t, err)
	assert.NotNil(t, response.Certifications)
	assert.NotNil(t, response.Education)
	assert.NotNil(t, response.Experiences)
	assert.NotNil(t, response.Readings)
}
