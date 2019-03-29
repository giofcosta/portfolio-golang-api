package models_test

import (
	"testing"

	"github.com/giofcosta/portfolio-golang-api/models"
	"github.com/stretchr/testify/assert"
)

func TestAboutMe_GetAll(t *testing.T) {
	model := &models.AboutMe{}

	data := model.GetAll()

	assert.NotNil(t, data)
}

func TestAboutMe_GetItems(t *testing.T) {
	model := &models.AboutMe{}

	items := model.GetItems()

	assert.NotNil(t, items)
}
