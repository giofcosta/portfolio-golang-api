package models

type Skill struct {
	Items     []Items     `json:"items"`
	Languages []Languages `json:"languages"`
	Tags      []string    `json:"tags"`
}
type Items struct {
	Title string `json:"title"`
	Stars int    `json:"stars"`
}
type Languages struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
