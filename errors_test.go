package uiiclient

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseError(t *testing.T) {
	assert := assert.New(t)

	t.Run("success", func(t *testing.T) {
		b, err := os.ReadFile("testdata/error.json")
		assert.NoError(err)

		actual := parseError(b)

		// we should def use RFC3339 instead of this custom shit
		expected := UIIError{
			Errors:     []string{"none of your networks has a valid internet connection"},
			Timestamp:  "2022-01-16T18:12:52+0000",
			StatusCode: http.StatusBadRequest,
			Instance:   "/images",
		}
		assert.Equal(expected, actual)
	})

	t.Run("error", func(t *testing.T) {
		actual := parseError([]byte("'"))
		assert.Error(actual)
	})
}
