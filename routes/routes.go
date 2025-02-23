package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"himtalks-backend/controllers"
	"himtalks-backend/middleware"
	"himtalks-backend/ws"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Endpoint untuk pesan anonim dan songfess
	messageController := &controllers.MessageController{DB: db}
	r.HandleFunc("/message", messageController.SendMessage).Methods("POST")

	songfessController := &controllers.SongfessController{DB: db}
	r.HandleFunc("/songfess", songfessController.SendSongfess).Methods("POST")

	// WebSocket
	r.HandleFunc("/ws", ws.HandleConnections)

	// OAuth Google
	adminController := &controllers.AdminController{}
	r.HandleFunc("/auth/google/login", adminController.Login).Methods("GET")
	r.HandleFunc("/auth/google/callback", adminController.Callback).Methods("GET")

	// Protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		email := r.Context().Value("email").(string)
		fmt.Fprintf(w, "Welcome, %s!", email)
	}).Methods("GET")

	// Tambahkan endpoint untuk mendapatkan daftar songfess dan pesan
	protected.HandleFunc("/songfess", songfessController.GetSongfessList).Methods("GET")
	protected.HandleFunc("/messages", messageController.GetMessageList).Methods("GET")

	return r
}
