package main

import (
	"log"
	"os"

	"himtalks-backend/config"
	"himtalks-backend/models"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load environment variables
	config.LoadEnv()

	// Connect to PostgreSQL
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Create admin table if it doesn't exist
	err = models.CreateTableAdmins(db)
	if err != nil {
		log.Fatal("Failed to create admins table:", err)
	}

	// Add admin from command line argument
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run add_admin.go <admin_email>")
	}

	adminEmail := os.Args[1]
	
	// Check if admin already exists
	isAdmin, err := models.IsAdmin(db, adminEmail)
	if err != nil {
		log.Fatal("Failed to check admin status:", err)
	}

	if isAdmin {
		log.Printf("Email %s is already an admin", adminEmail)
		return
	}

	// Add admin
	err = models.InsertAdmin(db, adminEmail)
	if err != nil {
		log.Fatal("Failed to add admin:", err)
	}

	log.Printf("Successfully added %s as admin", adminEmail)
}
