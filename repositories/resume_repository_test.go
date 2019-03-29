package repositories_test

import (
	"testing"

	"github.com/giofcosta/portfolio-golang-api/repositories"
	"github.com/stretchr/testify/assert"
)

var resumeRepository repositories.ResumeRepository

func init() {
	resumeRepository = repositories.NewResumeRepository()
}

func TestResume_GetAll(t *testing.T) {

	data := resumeRepository.GetAll()

	assert.NotNil(t, data)
}

func TestResume_GetCertifications(t *testing.T) {

	items := resumeRepository.GetCertifications()

	assert.NotNil(t, items)
}

func TestResume_GetEducation(t *testing.T) {

	items := resumeRepository.GetEducation()

	assert.NotNil(t, items)
}

func TestResume_GetExperiences(t *testing.T) {

	items := resumeRepository.GetExperiences()

	assert.NotNil(t, items)
}
func TestResume_GetReadings(t *testing.T) {

	items := resumeRepository.GetReadings()

	assert.NotNil(t, items)
}
