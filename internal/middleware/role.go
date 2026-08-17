package middleware

import "net/http"

func RequireRole(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole, ok := r.Context().Value(roleKey).(string)

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
