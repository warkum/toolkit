package gqlclient

import "github.com/machinebox/graphql"

// setHeaders set headers key and value to request headers
func setHeaders(req *graphql.Request, headers map[string]string) {
	if req == nil {
		return
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}
}

// setVariables set variables to graphql request
func setVariables(req *graphql.Request, variables map[string]interface{}) {
	if req == nil {
		return
	}

	for key, value := range variables {
		req.Var(key, value)
	}
}
