package requests

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	github "github.com/Defl8/what-is-wyatt-doing/internal/github"
)

type RequestType int

const (
	GET RequestType = iota
	POST
	PUT
	DELETE
)

type RequestHandler struct {
	BaseURL string
}

func NewRequestHandler(endpoint *string) *RequestHandler {

	if endpoint == nil {
		return &RequestHandler{
			BaseURL: "",
		}
	}

	return &RequestHandler{
		BaseURL: *endpoint,
	}
}

func (rH RequestHandler) GetPublicUserEvents(user string) (*[]github.Event, error) {
	rH.BaseURL = "https://api.github.com/"

	headers := map[string]string{
		"Accept": "application/vnd.github+json",
	}

	resp, err := rH.MakeRequest(GET, "users/Defl8/events", headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var events []github.Event
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, err
	}

	return &events, nil
}

func (rH RequestHandler) MakeRequest(reqType RequestType, endpoint string, headers map[string]string) (*http.Response, error) {
	fullURL := rH.BaseURL + endpoint

	// Validate URL before trying to make the request
	validURL, err := url.ParseRequestURI(fullURL)
	if err != nil {
		return nil, err
	}

	switch reqType {
	case GET:
		resp, err := rH.MakeGetRequest(validURL.String(), headers)
		if err != nil {
			return nil, err
		}
		return resp, nil

	default: // This shouldn't happen but just to be safe
		return nil, errors.New("Invalid request type")
	}
}

func (rH RequestHandler) MakeGetRequest(url string, headers map[string]string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
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
