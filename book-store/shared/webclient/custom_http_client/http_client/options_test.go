package http_client

import (
	"net/http"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestOptionsAreSet(t *testing.T) {

	httpTimeout := 25 * time.Millisecond

	client := &http.Client{Timeout: 25 * time.Millisecond}

	c := NewClient(
		WithHTTPClient(client),
		WithHTTPTimeout(httpTimeout),
	)

	assert.Equal(t, client, c.client)
	assert.Equal(t, httpTimeout, c.timeout)

}

func TestOptionsHaveDefaults(t *testing.T) {
	httpTimeout := 30 * time.Second
	http.DefaultClient.Timeout = httpTimeout
	c := NewClient()

	assert.Equal(t, http.DefaultClient, c.client)
	assert.Equal(t, httpTimeout, c.timeout)
}
