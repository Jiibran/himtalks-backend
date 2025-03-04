package models

import (
	"database/sql"
	"time"
)

type Message struct {
	ID            int       `json:"id"`
	Content       string    `json:"content"`
	SenderName    string    `json:"sender_name"`
	RecipientName string    `json:"recipient_name"`
	Category      string    `json:"category"` // "kritik" atau "saran"
	CreatedAt     time.Time `json:"created_at"`
}

// CreateTableMessages membuat tabel messages di PostgreSQL
func CreateTableMessages(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		sender_name VARCHAR(50),
		recipient_name VARCHAR(50),
		category VARCHAR(20) CHECK (category IN ('kritik', 'saran')),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}
