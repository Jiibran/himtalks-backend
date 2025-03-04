package models

import (
	"database/sql"
	"strconv"
	"time"
)

type Songfess struct {
	ID            int       `json:"id"`
	Content       string    `json:"content"`
	SongID        string    `json:"song_id"`
	SongTitle     string    `json:"song_title"`
	Artist        string    `json:"artist"`
	AlbumArt      string    `json:"album_art"`
	StartTime     int       `json:"start_time"`
	EndTime       int       `json:"end_time"`
	SenderName    string    `json:"sender_name"`
	RecipientName string    `json:"recipient_name"`
	CreatedAt     time.Time `json:"created_at"`
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
		sender_name VARCHAR(50),
		recipient_name VARCHAR(50),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}

// Tambahkan fungsi untuk mengambil batas karakter pesan
func GetMessageCharLimit(db *sql.DB) (int, error) {
	var stored string
	err := db.QueryRow("SELECT value FROM configs WHERE key='message_char_limit'").Scan(&stored)
	if err != nil {
		return 100, nil // Default 100 karakter jika belum diatur
	}

	charLimit, convErr := strconv.Atoi(stored)
	if convErr != nil {
		return 100, nil
	}
	return charLimit, nil
}
