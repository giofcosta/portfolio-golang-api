package models

//Resume defines the model for resume informations
type Resume struct {
	Experiences    []Experiences    `json:"experiences"`
	Education      []Education      `json:"education"`
	Certifications []Certifications `json:"certifications"`
	Readings       []Readings       `json:"readings"`
}

type Experiences struct {
	Title       string `json:"title"`
	Date        string `json:"date"`
	Rule        string `json:"rule"`
	Description string `json:"description"`
}
type Education struct {
	Title string `json:"title"`
	Rule  string `json:"rule"`
	Date  string `json:"date"`
}
type Certifications struct {
	Title       string `json:"title"`
	Rule        string `json:"rule"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
type Readings struct {
	Title       string `json:"title"`
	Rule        string `json:"rule"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
