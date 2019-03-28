package models

type AboutMe struct {
	Items []string `json:"items"`
}

func (am *AboutMe) GetAll() *AboutMe {
	return &AboutMe{
		Items: am.GetItems(),
	}
}

func (am *AboutMe) GetItems() []string {
	return []string{
		"Software engineer with 15+ years of experience as a developer.",
		"Developed several robust systems with high scalability.",
		"Worked on the project that changed the logistic for the largest e-commerce company in Latin America.",
		"Fast and creative thinker.",
		"Focused on solving the most difficult problems.",
	}
}
