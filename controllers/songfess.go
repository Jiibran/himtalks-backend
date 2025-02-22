package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"himtalks-backend/models"
)

type SongfessController struct {
	DB *sql.DB
}

// SendSongfess menangani pengiriman songfess
func (sc *SongfessController) SendSongfess(w http.ResponseWriter, r *http.Request) {
	var songfess models.Songfess
	err := json.NewDecoder(r.Body).Decode(&songfess)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Simpan songfess ke database
	query := `INSERT INTO songfess (content, song_id) VALUES ($1, $2) RETURNING id, created_at`
	err = sc.DB.QueryRow(query, songfess.Content, songfess.SongID).Scan(&songfess.ID, &songfess.CreatedAt)
	if err != nil {
		http.Error(w, "Failed to save songfess", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(songfess)
}
