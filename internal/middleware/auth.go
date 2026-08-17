package middleware

import (
	"context"
	"net/http"
	"prueba/internal/auth"
	"prueba/internal/requestcontext"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		token, err := auth.VerifyToken(parts[1])

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(w, "Invalid claims", http.StatusUnauthorized)
			return
		}

		role, ok := claims["role"].(string)

		if !ok {
			http.Error(w, "Invalid role", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"].(float64)

		if !ok {
			http.Error(w, "Invalid user ID", http.StatusUnauthorized)
			return
		}

		ctx := requestcontext.WithUserID(r.Context(), int(userID))
		ctx = context.WithValue(ctx, requestcontext.RoleKey, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
