package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// GqlTimeout is a middleware to add http.TimeoutHandler
// will return 503 http error code & graphQL error response body
func GqlTimeout(timeout time.Duration) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		resp, _ := json.Marshal(graphql.ErrorResponse(context.TODO(), "Timeout"))
		return http.TimeoutHandler(next, timeout, string(resp))
	}
}
