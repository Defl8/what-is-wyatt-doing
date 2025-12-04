package requests

import "net/http"

type RequestType int

const (
	GET RequestType = iota
	POST
	PUT
	DELETE
)

type RequestHandler struct {
	Endpoint string
}

func NewRequestHandler(endpoint string) *RequestHandler {
	return &RequestHandler{
		Endpoint: endpoint,
	}
}

func (rH RequestHandler) MakeGetRequest(headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", rH.Endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Build headers for request
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
