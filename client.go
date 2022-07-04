package uiiclient

import (
	"net/http"
	"net/url"
)

type UIIClient struct {
	token  string
	url    *url.URL
	client *http.Client
}

// NewClient creates a new UII api client.
func NewClient(apiToken string) (*UIIClient, error) {
	uri, err := url.ParseRequestURI("https://api.virtomize.com/uii")
	if err != nil {
		return nil, err
	}

	return &UIIClient{
		token:  apiToken,
		url:    uri,
		client: &http.Client{Transport: http.DefaultTransport},
	}, nil
}
