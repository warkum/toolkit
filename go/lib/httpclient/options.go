package httpclient

import "time"

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.HttpCli.Timeout = timeout
	}
}
