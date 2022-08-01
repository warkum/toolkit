package authctx

import (
	"context"
	"time"
)

type (
	UUID       string
	AUTH_LEVEL string
)

func (l AUTH_LEVEL) IsAuthorizedAs(minLevel AUTH_LEVEL) bool {
	rank, okRank := AUTH_RANK[l]
	minRank, okMinRank := AUTH_RANK[minLevel]
	if !okRank || !okMinRank {
		return false
	}

	return rank >= minRank
}

const (
	AUTH_SUPER AUTH_LEVEL = "SUPERADMIN"
	AUTH_ADMIN AUTH_LEVEL = "ADMINISTRATOR"
	AUTH_OPS   AUTH_LEVEL = "OPS"
	AUTH_BUYER AUTH_LEVEL = "BUYER"
)

var AUTH_RANK = map[AUTH_LEVEL]int{
	AUTH_BUYER: 0,
	AUTH_OPS:   1,
	AUTH_ADMIN: 2,
	AUTH_SUPER: 3,
}

const (
	UserIDHeaderKey       = "x-user-id"
	RoleHeaderKey         = "x-roles"
	HasuraUserIDHeaderKey = "x-hasura-user-id"
	HasuraRoleHeaderKey   = "x-hasura-role"
	UserLevelHeaderKey    = "x-user-level"
	AppPlatformHeaderKey  = "warkum-app-platform"
	AppVersionHeaderKey   = "warkum-app-version"
)

type Auth struct {
	UserID     UUID
	Msisdn     string
	Level      AUTH_LEVEL
	Expiry     time.Duration
	Platform   string
	AppVersion string
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

func GetPlatformInContext(ctx context.Context) string {
	value, _ := ctx.Value(XUserIDCtxKey).(*Auth)
	var platform string
	if value != nil {
		platform = string(value.Platform)
	}
	return platform
}

func GetAppVersionInContext(ctx context.Context) string {
	value, _ := ctx.Value(XUserIDCtxKey).(*Auth)
	var version string
	if value != nil {
		version = string(value.AppVersion)
	}
	return version
}
