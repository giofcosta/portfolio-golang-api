package repositories

import "github.com/giofcosta/portfolio-golang-api/models"

type AboutMeRepository interface {
	GetAll() *models.AboutMe
	GetItems() []string
}

type aboutMeRepository struct {
}

func NewAboutMeRepository() *aboutMeRepository {
	return &aboutMeRepository{}
}

func (a *aboutMeRepository) GetAll() *models.AboutMe {
	return &models.AboutMe{
		Items: a.GetItems(),
	}
}

func (a *aboutMeRepository) GetItems() []string {
	return []string{
		"Software engineer with 15+ years of experience as a developer.",
		"Developed several robust systems with high scalability.",
		"Worked on the project that changed the logistic for the largest e-commerce company in Latin America.",
		"Fast and creative thinker.",
		"Focused on solving the most difficult problems.",
	}
}
