package uiiclient

import "net/http"

const (
	DefaultUserAgent = "cloud-api-client-go-v1.0.0"
)

func (c *UIIClient) setDefaultHeader(r *http.Request) {
	r.Header.Add("Authorization", "Bearer "+c.token)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("User-Agent", DefaultUserAgent)
}
