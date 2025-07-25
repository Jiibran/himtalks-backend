package main

import (
	"log"

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

	log.Println("=== Database Debug Information ===")

	// Check if tables exist
	tables := []string{"messages", "songfess", "admins", "configs", "blacklist_words"}
	
	for _, table := range tables {
		var exists bool
		query := `SELECT EXISTS (
			SELECT FROM information_schema.tables 
			WHERE table_name = $1
		)`
		err := db.QueryRow(query, table).Scan(&exists)
		if err != nil {
			log.Printf("Error checking table %s: %v", table, err)
		} else {
			log.Printf("Table %s exists: %t", table, exists)
		}
	}

	// Check admin users
	log.Println("\n=== Admin Users ===")
	rows, err := db.Query("SELECT email FROM admins")
	if err != nil {
		log.Printf("Error querying admins: %v", err)
	} else {
		defer rows.Close()
		count := 0
		for rows.Next() {
			var email string
			err := rows.Scan(&email)
			if err != nil {
				log.Printf("Error scanning admin: %v", err)
			} else {
				log.Printf("Admin: %s", email)
				count++
			}
		}
		log.Printf("Total admins: %d", count)
	}

	// Check messages count
	var messageCount int
	err = db.QueryRow("SELECT COUNT(*) FROM messages").Scan(&messageCount)
	if err != nil {
		log.Printf("Error counting messages: %v", err)
	} else {
		log.Printf("Total messages: %d", messageCount)
	}

	// Check songfess count
	var songfessCount int
	err = db.QueryRow("SELECT COUNT(*) FROM songfess").Scan(&songfessCount)
	if err != nil {
		log.Printf("Error counting songfess: %v", err)
	} else {
		log.Printf("Total songfess: %d", songfessCount)
	}

	// Check configs
	log.Println("\n=== System Configs ===")
	configs, err := models.GetAllConfigs(db)
	if err != nil {
		log.Printf("Error getting configs: %v", err)
	} else {
		for key, value := range configs {
			log.Printf("Config %s: %s", key, value)
		}
	}

	log.Println("\n=== Debug completed ===")
}
