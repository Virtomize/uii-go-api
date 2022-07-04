package uiiclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperatingSystems(t *testing.T) {
	assert := assert.New(t)

	t.Run("success", func(t *testing.T) {
		c, err := NewClient("my-token")
		assert.NoError(err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := os.ReadFile("testdata/os.json")
			assert.NoError(err)
			_, _ = w.Write(b)
		}))

		c.url, err = url.Parse(server.URL)
		assert.NoError(err)

		actual, err := c.OperatingSystems()
		assert.NoError(err)

		expected := []OS{}
		b, err := os.ReadFile("testdata/TestOperatingSystems_success.golden")
		assert.NoError(err)

		err = json.Unmarshal(b, &expected)
		assert.NoError(err)

		assert.Equal(expected, actual)
	})

	t.Run("failure-api-error", func(t *testing.T) {
		c, err := NewClient("my-token")
		assert.NoError(err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := os.ReadFile("testdata/error.json")
			assert.NoError(err)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(b)
		}))

		c.url, err = url.Parse(server.URL)
		assert.NoError(err)

		_, err = c.OperatingSystems()
		assert.Error(err)
	})

	t.Run("failure-unknown-error", func(t *testing.T) {
		c, err := NewClient("my-token")
		assert.NoError(err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("some error"))
		}))

		c.url, err = url.Parse(server.URL)
		assert.NoError(err)

		_, err = c.OperatingSystems()
		assert.Error(err)
	})
}
