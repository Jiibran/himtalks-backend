package middleware

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"himtalks-backend/models"
	"himtalks-backend/utils"
)

// Gunakan string konstan untuk konsistensi
const emailKey = "email"

func AuthMiddlewareAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Coba ambil token dari cookie terlebih dahulu
		cookie, err := r.Cookie("jwt")
		var tokenString string

		if err == nil {
			// Token ditemukan di cookie
			tokenString = cookie.Value
		} else {
			// Fallback ke header Authorization (untuk kompatibilitas)
			tokenString = r.Header.Get("Authorization")
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		// Cek jika tidak ada token
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Validasi token
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Simpan email di context
		ctx := context.WithValue(r.Context(), emailKey, claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CheckIsAdmin(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Gunakan key yang sama dengan di AuthMiddleware
			email, ok := r.Context().Value(emailKey).(string)
			if !ok {
				http.Error(w, "Unauthorized: missing email", http.StatusUnauthorized)
				return
			}

			isAdmin, err := models.IsAdmin(db, email)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if !isAdmin {
				http.Error(w, "Forbidden: not an admin", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
