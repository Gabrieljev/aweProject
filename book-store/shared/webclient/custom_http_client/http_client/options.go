package http_client

import (
	"net/http"
	"time"
)

// Option represents the client options
type Option func(*Client)

// WithHTTPTimeout sets timeout
func WithHTTPTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}

// WithHTTPClient sets a custom http client
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

// WithReuse sets reuse
func WithReuse(isReuse bool) Option {
	return func(c *Client) {
		c.isReuse = isReuse
	}
}
