package authctx

import (
	"context"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xUserID := r.Header.Get(UserIDHeaderKey)
		log.Debug("[DEBUG] userID ", xUserID)

		if xUserID == "" {
			next.ServeHTTP(w, r)
			return
		}

		// assume roles value like this "admin,ops"
		xRoles := r.Header.Get(RoleHeaderKey)
		log.Debug("[DEBUG] roles ", xRoles)

		var level = AUTH_BUYER
		if isContainAdmin := strings.Contains(xRoles, "ops"); isContainAdmin {
			level = AUTH_OPS
		}
		if isContainAdmin := strings.Contains(xRoles, "admin"); isContainAdmin {
			level = AUTH_ADMIN
		}

		userXUserID := Auth{
			UserID: UUID(xUserID),
			Level:  level,
		}

		// put it in context
		ctx := context.WithValue(r.Context(), XUserIDCtxKey, &userXUserID)

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
