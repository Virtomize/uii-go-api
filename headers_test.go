package uiiclient

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setDefaultHeader(t *testing.T) {
	assert := assert.New(t)

	c, err := NewClient("my-token")
	assert.NoError(err)

	r, err := http.NewRequest(http.MethodGet, "test", http.NoBody)
	assert.NoError(err)
	c.setDefaultHeader(r)

	assert.Equal("Bearer my-token", r.Header.Get("Authorization"))
	assert.Equal("application/json", r.Header.Get("Content-Type"))
	assert.Equal(DefaultUserAgent, r.Header.Get("User-Agent"))
}
