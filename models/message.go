package models

import (
	"database/sql"
	"time"
)

type Message struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateTableMessages membuat tabel messages di PostgreSQL
func CreateTableMessages(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}
