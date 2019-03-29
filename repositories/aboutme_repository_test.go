package repositories_test

import (
	"testing"

	"github.com/giofcosta/portfolio-golang-api/repositories"
	"github.com/stretchr/testify/assert"
)

var aboutmeRepository repositories.AboutMeRepository

func init() {
	aboutmeRepository = repositories.NewAboutMeRepository()
}

func TestAboutMe_GetAll(t *testing.T) {

	data := aboutmeRepository.GetAll()

	assert.NotNil(t, data)
}

func TestAboutMe_GetItems(t *testing.T) {

	items := aboutmeRepository.GetItems()

	assert.NotNil(t, items)
}
