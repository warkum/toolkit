package authctx

import (
	"context"
	"time"
)

type (
	UUID       string
	AUTH_LEVEL string
)

const (
	AUTH_SUPER AUTH_LEVEL = "SUPERADMIN"
	AUTH_ADMIN AUTH_LEVEL = "ADMINISTRATOR"
	AUTH_OPS   AUTH_LEVEL = "OPS"
	AUTH_BUYER AUTH_LEVEL = "BUYER"
)

const (
	UserIDHeaderKey = "x-user-id"
	RoleHeaderKey   = "x-roles"
)

type Auth struct {
	UserID UUID
	Msisdn string
	Level  AUTH_LEVEL
	Expiry time.Duration
}

type ContextKey struct {
	Name string
}

var (
	XUserIDCtxKey = &ContextKey{Name: "x-user-id"}
)

// GetUserIDInContext finds the userID in the context. REQUIRES Middleware to have run.
func GetUserIDInContext(ctx context.Context) string {
	value, _ := ctx.Value(XUserIDCtxKey).(*Auth)
	var userID string
	if value != nil {
		userID = string(value.UserID)
	}
	return userID
}

func GetUserLevelInContext(ctx context.Context) AUTH_LEVEL {
	value, _ := ctx.Value(XUserIDCtxKey).(*Auth)
	var userLevel AUTH_LEVEL
	if value != nil {
		userLevel = AUTH_LEVEL(value.Level)
	}
	return userLevel
}

func GetAuthInContext(ctx context.Context) *Auth {
	value, _ := ctx.Value(XUserIDCtxKey).(*Auth)
	return value
}
