package unattended_install_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// interface for a http client (used for testing)
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func mergeToJsonByteSlice(argument1 interface{}, argument2 interface{}) ([]byte, error) {
	arg1Map, err := toStringMap(argument1)
	if err != nil {
		return nil, err
	}

	arg2Map, err := toStringMap(argument2)
	if err != nil {
		return nil, err
	}

	for k, v := range arg2Map {
		arg1Map[k] = v
	}

	return json.Marshal(arg1Map)
}

func toStringMap(argument interface{}) (map[string]interface{}, error) {
	var argumentAsMap map[string]interface{}
	jsonString, err := json.Marshal(argument)
	if err != nil {
		return argumentAsMap, err
	}

	err = json.Unmarshal(jsonString, &argumentAsMap)
	if err != nil {
		return nil, err
	}
	return argumentAsMap, err
}

func (c VirtomizeCloudApiClient) requestStringResult(method string, uri string, jsonBody []byte) (string, error) {
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+c.token)
	addSharedHeaders(&req.Header)

	if string(jsonBody) != "" {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		returnedError := parseError(b)
		if returnedError != "" {
			return "", fmt.Errorf(returnedError)
		}

		return "", fmt.Errorf("request was not successful and returned status %d and body \"%s\"", resp.StatusCode, string(b))
	}

	return string(b), nil
}

func addSharedHeaders(h *http.Header) {
	h.Add("User-Agent", "cloud-api-client-go-1.0.0")
}

func parseError(b []byte) string {
	var msgErr struct {
		Error []string `json:"errors"`
		Err   []error
	}
	jsonError := json.Unmarshal(b, &msgErr)
	if jsonError != nil {
		return ""
	}
	for _, e := range msgErr.Error {
		return e
	}

	return ""
}
