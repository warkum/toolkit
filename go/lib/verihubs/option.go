package verihubs

import (
	"net/http"
	"time"
)

// WithTimeout configure the client to have the specified http timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.http.Timeout = timeout
	}
}

// WithTransport configure the client to use the specified http transport
func WithTransport(transport http.RoundTripper) Option {
	return func(c *Client) {
		c.http.Transport = transport
	}
}

// WithCookieJar configure the client to use the specified cookie jar
func WithCookieJar(jar http.CookieJar) Option {
	return func(c *Client) {
		c.http.Jar = jar
	}
}
