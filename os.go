package uiiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// OperatingSystems returns available operating systems.
func (c *UIIClient) OperatingSystems() ([]OS, error) {
	uri := fmt.Sprintf("%s/oslist", c.url.String())

	res, err := c.defaultRequest(http.MethodGet, uri, []byte{})
	if err != nil {
		return nil, err
	}

	var response struct {
		Embedded []OS `json:"_embedded"`
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response.Embedded, nil
}
