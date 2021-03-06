package uiiclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	assert := assert.New(t)

	args := BuildArgs{
		Distribution: "UnitTest",
		Version:      "1",
		Hostname:     "Horst",
		Networks:     []NetworkArgs{{DHCP: true}},
	}
	opts := BuildOpts{}

	t.Run("success", func(t *testing.T) {
		c, err := NewClient("my-token")
		assert.NoError(err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := os.ReadFile("testdata/test.iso")
			assert.NoError(err)
			_, _ = w.Write(b)
		}))

		c.url, err = url.Parse(server.URL)
		assert.NoError(err)

		err = c.Build("testdata/build-test.iso", args, opts)
		assert.NoError(err)

		b, err := os.ReadFile("testdata/build-test.iso")
		assert.NoError(err)
		assert.Equal("test\n", string(b))
	})

	t.Run("failure-api-error", func(t *testing.T) {
		c, err := NewClient("my-token")
		assert.NoError(err)
		b, err := os.ReadFile("testdata/error.json")
		assert.NoError(err)

		expected := parseError(b)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(b)
		}))

		c.url, err = url.Parse(server.URL)
		assert.NoError(err)

		err = c.Build("testdata/build-test.iso", args, opts)
		assert.Error(err)
		assert.Equal(fmt.Errorf("unexpected status code 400: %w", expected), err)
	})

	t.Run("failure-unknown-error", func(t *testing.T) {
		c, err := NewClient("my-token")
		assert.NoError(err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("some error"))
		}))

		c.url, err = url.Parse(server.URL)
		assert.NoError(err)

		err = c.Build("testdata/build-test.iso", args, opts)
		assert.Error(err)
	})
}
