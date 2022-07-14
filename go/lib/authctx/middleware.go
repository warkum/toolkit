package authctx

import (
	"context"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var auth Auth

		auth.Platform = r.Header.Get(AppPlatformHeaderKey)
		auth.AppVersion = r.Header.Get(AppVersionHeaderKey)

		xUserID := r.Header.Get(UserIDHeaderKey)
		if xUserID == "" {
			xUserID = r.Header.Get(HasuraUserIDHeaderKey)
		}
		log.Debug("[DEBUG] userID ", xUserID)

		// user ID headers not found
		if xUserID == "" {
			ctx := context.WithValue(r.Context(), XUserIDCtxKey, &auth)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		}
		auth.UserID = UUID(xUserID)

		xUserLevel := r.Header.Get(UserLevelHeaderKey)
		if xUserLevel == "" {
			// assume roles value like this "admin,ops"
			xRoles := r.Header.Get(RoleHeaderKey)
			if xRoles == "" {
				xRoles = r.Header.Get(HasuraRoleHeaderKey)
			}
			log.Debug("[DEBUG] roles ", xRoles)

			auth.Level = AUTH_BUYER
			if isContainAdmin := strings.Contains(xRoles, "ops"); isContainAdmin {
				auth.Level = AUTH_OPS
			}
			if isContainAdmin := strings.Contains(xRoles, "admin"); isContainAdmin {
				auth.Level = AUTH_ADMIN
			}
			if isContainAdmin := strings.Contains(xRoles, "super"); isContainAdmin {
				auth.Level = AUTH_SUPER
			}
		} else {
			auth.Level = AUTH_LEVEL(xUserLevel)
		}

		// put it in context
		ctx := context.WithValue(r.Context(), XUserIDCtxKey, &auth)

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
