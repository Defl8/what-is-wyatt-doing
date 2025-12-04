package github

type Payload struct {
	Head   string `json:"head"`
	Before string `json:"before"`
}
