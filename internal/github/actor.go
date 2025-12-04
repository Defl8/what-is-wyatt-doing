package github

type Actor struct {
	ID           string    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	URL          string `json:"url"`
}
