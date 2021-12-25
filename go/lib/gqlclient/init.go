package gqlclient

import "github.com/machinebox/graphql"

// Client define dependencies data for gql client
type Client struct {
	client  *graphql.Client
	headers map[string]string
}

func New(cfg Config) *Client {
	client := graphql.NewClient(cfg.Address)

	return &Client{
		client:  client,
		headers: cfg.Headers,
	}
}
