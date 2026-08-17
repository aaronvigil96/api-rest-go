package middleware

import (
	"net/http"
	"prueba/internal/requestcontext"
)

func RequireRole(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole, ok := requestcontext.GetRole(r.Context())

			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if userRole != role {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
