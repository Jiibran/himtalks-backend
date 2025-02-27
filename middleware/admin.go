package middleware

import (
	"net/http"

	"database/sql"

	"himtalks-backend/models"
)

func CheckIsAdmin(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			email, _ := r.Context().Value("email").(string)
			ok, err := models.IsAdmin(db, email)
			if err != nil || !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
