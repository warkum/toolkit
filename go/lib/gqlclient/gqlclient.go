package gqlclient

import (
	"context"

	"github.com/machinebox/graphql"
)

// Run execute query or mutation by defined request
func (p *Client) Run(ctx context.Context, req Request, result interface{}) error {
	newReq := graphql.NewRequest(req.Message)

	if len(p.headers) > 0 {
		setHeaders(newReq, p.headers)
	}

	if len(req.Headers) > 0 {
		setHeaders(newReq, req.Headers)
	}

	if len(req.Variables) > 0 {
		setVariables(newReq, req.Variables)
	}

	return p.client.Run(ctx, newReq, result)
}
