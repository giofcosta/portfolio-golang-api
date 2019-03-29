package models_test

import (
	"testing"

	"github.com/giofcosta/portfolio-golang-api/models"
	"github.com/stretchr/testify/assert"
)

func TestResume_GetAll(t *testing.T) {
	model := &models.Resume{}

	data := model.GetAll()

	assert.NotNil(t, data)
}

func TestResume_GetCertifications(t *testing.T) {
	model := &models.Resume{}

	items := model.GetCertifications()

	assert.NotNil(t, items)
}

func TestResume_GetEducation(t *testing.T) {
	model := &models.Resume{}

	items := model.GetEducation()

	assert.NotNil(t, items)
}

func TestResume_GetExperiences(t *testing.T) {
	model := &models.Resume{}

	items := model.GetExperiences()

	assert.NotNil(t, items)
}
func TestResume_GetReadings(t *testing.T) {
	model := &models.Resume{}

	items := model.GetReadings()

	assert.NotNil(t, items)
}
