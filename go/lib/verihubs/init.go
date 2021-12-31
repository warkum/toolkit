package verihubs

import (
	"net/http"
)

// Client dependencies for verihubs client
type Client struct {
	http   *http.Client
	host   string
	apiKey string
	appID  string
}

// Option define an option for the client
type Option func(*Client)

// New init new client for verihubs
func New(config Config, opts ...Option) *Client {
	host := V1Host
	if config.Host != "" {
		host = config.Host
	}

	c := &Client{
		http: &http.Client{
			Timeout: config.Timeout,
		},
		host:   host,
		apiKey: config.APIKey,
		appID:  config.AppID,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
