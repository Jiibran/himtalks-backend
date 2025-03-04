package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"himtalks-backend/models"
	"himtalks-backend/ws"
)

type SongfessController struct {
	DB *sql.DB
}

// SendSongfess menangani pengiriman songfess
func (sc *SongfessController) SendSongfess(w http.ResponseWriter, r *http.Request) {
	log.Println("Received POST request to /songfess")
	var songfess models.Songfess
	err := json.NewDecoder(r.Body).Decode(&songfess)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validasi panjang pesan berdasarkan konfigurasi
	charLimit, _ := models.GetMessageCharLimit(sc.DB)
	if len(songfess.Content) > charLimit {
		http.Error(w, fmt.Sprintf("Message content exceeds character limit of %d", charLimit), http.StatusBadRequest)
		return
	}

	// Simpan songfess ke database
	query := `INSERT INTO songfess (content, song_id, song_title, artist, album_art, start_time, end_time, sender_name, recipient_name) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
              RETURNING id, created_at`

	err = sc.DB.QueryRow(
		query,
		songfess.Content,
		songfess.SongID,
		songfess.SongTitle,
		songfess.Artist,
		songfess.AlbumArt,
		songfess.StartTime,
		songfess.EndTime,
		songfess.SenderName,
		songfess.RecipientName,
	).Scan(&songfess.ID, &songfess.CreatedAt)

	if err != nil {
		http.Error(w, "Failed to save songfess: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Broadcast songfess to WebSocket clients
	ws.BroadcastMessage(ws.Message{
		Type: "songfess",
		Data: songfess,
	})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(songfess)
	log.Printf("Received songfess: %+v", songfess)
}

// GetSongfessList mengembalikan daftar songfess
func (sc *SongfessController) GetSongfessList(w http.ResponseWriter, r *http.Request) {
	rows, err := sc.DB.Query(
		"SELECT id, content, song_id, song_title, artist, album_art, start_time, end_time, sender_name, recipient_name, created_at FROM songfess")
	if err != nil {
		http.Error(w, "Failed to fetch songfess", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var songfessList []models.Songfess
	for rows.Next() {
		var songfess models.Songfess
		err := rows.Scan(
			&songfess.ID,
			&songfess.Content,
			&songfess.SongID,
			&songfess.SongTitle,
			&songfess.Artist,
			&songfess.AlbumArt,
			&songfess.StartTime,
			&songfess.EndTime,
			&songfess.SenderName,
			&songfess.RecipientName,
			&songfess.CreatedAt)
		if err != nil {
			http.Error(w, "Failed to scan songfess", http.StatusInternalServerError)
			return
		}
		songfessList = append(songfessList, songfess)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songfessList)
}

// Hanya menampilkan data >= cutoff
func (sc *SongfessController) GetSongfessListWithCutoff(w http.ResponseWriter, r *http.Request, cutoff time.Time) {
	rows, err := sc.DB.Query(`
        SELECT id, content, song_id, song_title, artist, album_art, start_time, end_time, sender_name, recipient_name, created_at 
        FROM songfess 
        WHERE created_at >= $1`, cutoff)
	if err != nil {
		http.Error(w, "Failed to fetch songfess", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var songfessList []models.Songfess
	for rows.Next() {
		var songfess models.Songfess
		err := rows.Scan(
			&songfess.ID,
			&songfess.Content,
			&songfess.SongID,
			&songfess.SongTitle,
			&songfess.Artist,
			&songfess.AlbumArt,
			&songfess.StartTime,
			&songfess.EndTime,
			&songfess.SenderName,
			&songfess.RecipientName,
			&songfess.CreatedAt)
		if err != nil {
			http.Error(w, "Failed to scan songfess", http.StatusInternalServerError)
			return
		}
		songfessList = append(songfessList, songfess)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songfessList)
}
