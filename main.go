package main

import (
	"log"
	"net/http"

	"himtalks-backend/config"
	"himtalks-backend/models"
	"himtalks-backend/routes"
	"himtalks-backend/ws"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to PostgreSQL
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Create tables if they don't exist
	err = models.CreateTableMessages(db)
	if err != nil {
		log.Fatal("Failed to create messages table:", err)
	}
	err = models.CreateTableSongfess(db)
	if err != nil {
		log.Fatal("Failed to create songfess table:", err)
	}

	// Setup routes
	r := routes.SetupRoutes(db)

	// WebSocket endpoint dengan CSP header
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Tambahkan CSP header untuk mengizinkan koneksi WebSocket
		w.Header().Set("Content-Security-Policy", "connect-src 'self' ws://localhost:8080")
		ws.HandleConnections(w, r)
	})

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
