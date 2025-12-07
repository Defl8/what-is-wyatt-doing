package github

type Payload struct {
	Head    string `json:"head"`
	Before  string `json:"before"`
	Ref     string `json:"ref"`
	RefType string `json:"ref_type"`
}
