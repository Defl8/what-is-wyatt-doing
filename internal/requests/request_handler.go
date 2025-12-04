package requests

type RequestHandler struct {
	Endpoint string
	Headers  map[string]any
}

func NewRequestHandler(endpoint string, headers map[string]any) *RequestHandler {
	return &RequestHandler{
		Endpoint: endpoint,
		Headers:  headers,
	}
}
