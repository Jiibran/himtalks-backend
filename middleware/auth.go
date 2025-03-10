package middleware

import (
	"context"
	"net/http"
	"strings"

	"himtalks-backend/utils"
)

type contextKey string

const emailContextKey contextKey = "email"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari header Authorization
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Hilangkan "Bearer " dari token
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Validasi token
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Simpan email di context dengan key yang sama
		ctx := context.WithValue(r.Context(), emailContextKey, claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
