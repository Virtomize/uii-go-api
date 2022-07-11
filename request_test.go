package uiiclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_defaultRequest(t *testing.T) {
	assert := assert.New(t)

	t.Run("success", func(t *testing.T) {
		c, err := NewClient("my-token")
		assert.NoError(err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("success"))
		}))

		b, err := c.defaultRequest(http.MethodGet, server.URL, nil)
		assert.NoError(err)

		assert.Equal("success", string(b))
	})

	t.Run("failure-no-body", func(t *testing.T) {
		c, err := NewClient("my-token")
		assert.NoError(err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		}))

		_, err = c.defaultRequest(http.MethodGet, server.URL, nil)
		assert.Error(err)
	})

	t.Run("failure-status-code", func(t *testing.T) {
		c, err := NewClient("my-token")
		assert.NoError(err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			b, err := json.Marshal(UIIError{Errors: []string{"some error"}})
			assert.NoError(err)
			_, _ = w.Write(b)
		}))

		_, err = c.defaultRequest(http.MethodGet, server.URL, nil)
		assert.Error(err)
	})
}
