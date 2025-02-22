package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"himtalks-backend/models"
)

type MessageController struct {
	DB *sql.DB
}

// SendMessage menangani pengiriman pesan anonim
func (mc *MessageController) SendMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Simpan pesan ke database
	query := `INSERT INTO messages (content) VALUES ($1) RETURNING id, created_at`
	err = mc.DB.QueryRow(query, message.Content).Scan(&message.ID, &message.CreatedAt)
	if err != nil {
		http.Error(w, "Failed to save message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}
