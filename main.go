package main

import (
	"log"
	"net/http"
	"os"

	"himtalks-backend/config"
	"himtalks-backend/models"
	"himtalks-backend/routes"
	"himtalks-backend/ws"

	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Akses environment variables
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := os.Getenv("GOOGLE_REDIRECT_URL")
	secretKey := os.Getenv("SECRET_KEY")

	log.Println("Client ID:", clientID)
	log.Println("Client Secret:", clientSecret)
	log.Println("Redirect URL:", redirectURL)
	log.Println("Secret Key:", secretKey)

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
