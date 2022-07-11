package uiiclient

import (
	"encoding/json"
)

type UIIError struct {
	Errors     []string `json:"errors"`
	Timestamp  string   `json:"timestamp"`
	StatusCode int      `json:"statuscode"`
	Instance   string   `json:"instance"`
}

// implement error interface
// returns a comma separated list of errors
func (e UIIError) Error() string {
	str := ""
	for k, v := range e.Errors {
		str += v
		if k > len(e.Errors)-1 {
			str += ", "
		}
	}
	return str
}

func parseError(b []byte) error {
	uiiErr := UIIError{}
	err := json.Unmarshal(b, &uiiErr)
	if err != nil {
		return err
	}
	return uiiErr
}
