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

// GetSongfessList mengembalikan daftar songfess
func (sc *SongfessController) GetSongfessList(w http.ResponseWriter, r *http.Request) {
	rows, err := sc.DB.Query("SELECT id, content, song_id, created_at FROM songfess")
	if err != nil {
		http.Error(w, "Failed to fetch songfess", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var songfessList []models.Songfess
	for rows.Next() {
		var songfess models.Songfess
		err := rows.Scan(&songfess.ID, &songfess.Content, &songfess.SongID, &songfess.CreatedAt)
		if err != nil {
			http.Error(w, "Failed to scan songfess", http.StatusInternalServerError)
			return
		}
		songfessList = append(songfessList, songfess)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songfessList)
}
