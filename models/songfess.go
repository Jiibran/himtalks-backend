package models

import (
	"database/sql"
	"time"
)

type Songfess struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	SongID    string    `json:"song_id"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateTableSongfess membuat tabel songfess di PostgreSQL
func CreateTableSongfess(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS songfess (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		song_id TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}
