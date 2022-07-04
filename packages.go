package uiiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Packages returns all available packages for a given OS.
func (c *UIIClient) Packages(args PackageArgs, opts PackageOpts) ([]string, error) {
	uri := fmt.Sprintf("%s/images", c.url.String())

	jsonPayload, err := json.Marshal(packageCombined{args, opts})
	if err != nil {
		return []string{}, err
	}

	res, err := c.defaultRequest(http.MethodPost, uri, jsonPayload)
	if err != nil {
		return []string{}, err
	}

	var response struct {
		Embedded PackageListResponse `json:"_embedded"`
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		return []string{}, err
	}

	return response.Embedded.Packages, nil
}
