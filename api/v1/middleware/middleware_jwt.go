package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gasBlar/GoGoManager/utils"
)

var whitelistPath = []string{
	"/api/v1/auth",
	"/metrics",
}

func isWhitelisted(path string) bool {
	for _, allowedPath := range whitelistPath {
		if strings.HasPrefix(path, allowedPath) {
			return true
		}
	}
	return false
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isWhitelisted(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" || !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		user, err := utils.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
