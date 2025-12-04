package github

type Actor struct {
	ID           int `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	URL          string `json:"url"`
}
