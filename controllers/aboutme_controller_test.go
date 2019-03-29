package controllers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/giofcosta/portfolio-golang-api/utils"

	"github.com/giofcosta/portfolio-golang-api/models"
	"github.com/giofcosta/portfolio-golang-api/server"
	"github.com/stretchr/testify/assert"
)

func TestAboutMeController_GET(t *testing.T) {
	router := server.NewRouter()

	w := utils.PerformRequest(router, "GET", "/aboutme")

	var response models.AboutMe
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response.Items)
}