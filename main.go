package main

import (
	"log"
	"net/http"
	"os"

	"himtalks-backend/config"
	"himtalks-backend/models"
	"himtalks-backend/routes"

	"himtalks-backend/ws"

	"github.com/gorilla/handlers"
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

	// Start WebSocket handler
	go ws.HandleMessages()

	// Middleware untuk menonaktifkan CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://himtalks.japaneast.cloudapp.azure.com", "https://himtalks-frontend.vercel.app"}), // Ganti * dengan domain FE
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(), // Mengaktifkan penggunaan credentials (cookies/session)
	)

	// Start server dengan middleware CORS
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}
