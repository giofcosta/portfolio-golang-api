package repositories_test

import (
	"testing"

	"github.com/giofcosta/portfolio-golang-api/repositories"

	"github.com/stretchr/testify/assert"
)

var skillRepository repositories.SkillRepository

func init() {
	skillRepository = repositories.NewSkillRepository()
}

func TestSkill_GetAll(t *testing.T) {
	data := skillRepository.GetAll()

	assert.NotNil(t, data)
}

func TestSkill_GetItems(t *testing.T) {
	items := skillRepository.GetItems()

	assert.NotNil(t, items)
}

func TestSkill_GetLanguages(t *testing.T) {

	items := skillRepository.GetLanguages()

	assert.NotNil(t, items)
}

func TestSkill_GetTags(t *testing.T) {

	items := skillRepository.GetTags()

	assert.NotNil(t, items)
}
