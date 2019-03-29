package models_test

import (
	"testing"

	"github.com/giofcosta/portfolio-golang-api/models"
	"github.com/stretchr/testify/assert"
)

func TestSkill_GetAll(t *testing.T) {
	model := &models.Skill{}

	data := model.GetAll()

	assert.NotNil(t, data)
}

func TestSkill_GetItems(t *testing.T) {
	model := &models.Skill{}

	items := model.GetItems()

	assert.NotNil(t, items)
}

func TestSkill_GetLanguages(t *testing.T) {
	model := &models.Skill{}

	items := model.GetLanguages()

	assert.NotNil(t, items)
}

func TestSkill_GetTags(t *testing.T) {
	model := &models.Skill{}

	items := model.GetTags()

	assert.NotNil(t, items)
}
