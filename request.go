package uiiclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func (c *UIIClient) defaultRequest(method string, uri string, jsonBody []byte) ([]byte, error) {
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	c.setDefaultHeader(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, parseError(b)
	}

	return b, nil
}
