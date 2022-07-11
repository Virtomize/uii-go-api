package uiiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// Build builds a ISO for a given operating system configuration.
func (c *UIIClient) Build(filePath string, args BuildArgs, opts BuildOpts) error {
	req, err := c.createBuildRequest(args, opts)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return parseBuildResult(resp, filePath)
}

func (c *UIIClient) createBuildRequest(args BuildArgs, opts BuildOpts) (*http.Request, error) {
	uri := fmt.Sprintf("%s/images", c.url.String())

	jsonString, err := json.Marshal(buildCombined{args, opts})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonString))
	if err != nil {
		return nil, err
	}
	c.setDefaultHeader(req)

	return req, nil
}

func parseBuildResult(resp *http.Response, filePath string) error {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		err := parseError(b)
		if err != nil {
			return fmt.Errorf("unexpected status code %d: %w", resp.StatusCode, err)
		}
	}

	// response is a byte stream
	var file *os.File
	file, err = os.Create(filepath.Clean(filePath))
	if err != nil {
		return err
	}

	_, err = file.Write(b)
	if err != nil {
		_ = os.Remove(filePath)
		return err
	}

	return file.Close()
}
