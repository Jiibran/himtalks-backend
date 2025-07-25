package middleware

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strings"

	"himtalks-backend/models"
	"himtalks-backend/utils"
)

// Gunakan string konstan untuk konsistensi
const emailKey = "email"

func AuthMiddlewareAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("AuthMiddlewareAdmin: processing request for %s", r.URL.Path)
		
		// Coba ambil token dari cookie terlebih dahulu
		cookie, err := r.Cookie("jwt")
		var tokenString string

		if err == nil {
			// Token ditemukan di cookie
			tokenString = cookie.Value
			log.Printf("AuthMiddlewareAdmin: JWT token found in cookie")
		} else {
			// Fallback ke header Authorization (untuk kompatibilitas)
			tokenString = r.Header.Get("Authorization")
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			if tokenString != "" {
				log.Printf("AuthMiddlewareAdmin: JWT token found in Authorization header")
			}
		}

		// Cek jika tidak ada token
		if tokenString == "" {
			log.Printf("AuthMiddlewareAdmin: no JWT token found for %s", r.URL.Path)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Validasi token
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			log.Printf("AuthMiddlewareAdmin: invalid JWT token for %s: %v", r.URL.Path, err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		log.Printf("AuthMiddlewareAdmin: valid JWT token for email: %s", claims.Email)

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
				log.Printf("CheckIsAdmin: missing email in context for %s", r.URL.Path)
				http.Error(w, "Unauthorized: missing email", http.StatusUnauthorized)
				return
			}

			log.Printf("CheckIsAdmin: checking admin status for email: %s", email)
			
			isAdmin, err := models.IsAdmin(db, email)
			if err != nil {
				log.Printf("CheckIsAdmin: database error checking admin status for %s: %v", email, err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if !isAdmin {
				log.Printf("CheckIsAdmin: user %s is not an admin", email)
				http.Error(w, "Forbidden: not an admin", http.StatusForbidden)
				return
			}

			log.Printf("CheckIsAdmin: user %s is admin, allowing access to %s", email, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}
}
