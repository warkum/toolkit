package httpclient

import (
	"net/http"
)

type Client struct {
	HttpCli http.Client
}

type Option func(*Client)

func New(opts ...Option) *Client {
	cli := &Client{}

	for _, opt := range opts {
		opt(cli)
	}

	return cli
}
