package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"himtalks-backend/controllers"
	"himtalks-backend/middleware"
	"himtalks-backend/models"
	"himtalks-backend/ws"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Tambahkan Logger middleware di awal
	r.Use(middleware.Logger)
	r.Use(middleware.CORS)
	// Endpoint message
	messageController := &controllers.MessageController{DB: db}
	r.HandleFunc("/message", messageController.SendMessage).Methods("POST")

	// Endpoint songfess
	songfessController := &controllers.SongfessController{DB: db}
	// Tambahkan endpoint POST
	r.HandleFunc("/songfess", songfessController.SendSongfess).Methods("POST")
	// Ambil hanya data songfess 7 hari
	r.HandleFunc("/songfess", func(w http.ResponseWriter, r *http.Request) {
		days, _ := models.GetSongfessDays(db)
		cutoff := time.Now().AddDate(0, 0, -days)
		songfessController.GetSongfessListWithCutoff(w, r, cutoff)
	}).Methods("GET")
	// Endpoint untuk mengambil songfess berdasarkan ID
	r.HandleFunc("/songfess/{id:[0-9]+}", songfessController.GetSongfessById).Methods("GET")

	// Websocket
	r.HandleFunc("/ws", ws.HandleConnections)

	// OAuth
	adminController := &controllers.AdminController{}
	r.HandleFunc("/auth/google/login", adminController.Login).Methods("GET")
	r.HandleFunc("/auth/google/callback", adminController.Callback).Methods("GET")

	// Protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware) // cek JWT
	protected.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		email := r.Context().Value("email").(string)
		fmt.Fprintf(w, "Welcome, %s!", email)
	}).Methods("GET")

	// Spotify API Routes
	spotifyController := controllers.NewSpotifyController()
	r.HandleFunc("/api/spotify/search", spotifyController.SearchTracks).Methods("GET")
	r.HandleFunc("/api/spotify/track", spotifyController.GetTrack).Methods("GET")

	// Tambahan route admin
	admin := protected.PathPrefix("/admin").Subrouter()
	adminHandler := &controllers.AdminHandler{DB: db}
	admin.Use(middleware.CheckIsAdmin(db)) // Cek di middleware
	admin.HandleFunc("/addAdmin", adminHandler.AddAdmin).Methods("POST")
	admin.HandleFunc("/configSongfessDays", adminHandler.UpdateSongfessDays).Methods("POST")
	admin.HandleFunc("/blacklist", adminHandler.AddBlacklistWord).Methods("POST")
	admin.HandleFunc("/songfessAll", songfessController.GetSongfessList).Methods("GET") // tanpa cutoff
	admin.HandleFunc("/messages", messageController.GetMessageList).Methods("GET")
	admin.HandleFunc("/message/delete", messageController.DeleteMessage).Methods("POST")

	return r
}
