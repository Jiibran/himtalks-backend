package models

import (
	"database/sql"
	"time"
)

type Songfess struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	SongID    string    `json:"song_id"`
	SongTitle string    `json:"song_title"`
	Artist    string    `json:"artist"`
	AlbumArt  string    `json:"album_art"`
	StartTime int       `json:"start_time"`
	EndTime   int       `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateTableSongfess membuat tabel songfess di PostgreSQL
func CreateTableSongfess(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS songfess (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		song_id TEXT NOT NULL,
		song_title TEXT,
		artist TEXT,
		album_art TEXT,
		start_time INTEGER DEFAULT 0,
		end_time INTEGER DEFAULT 30,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}
