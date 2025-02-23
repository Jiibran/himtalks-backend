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

// GetMessageList mengembalikan daftar pesan
func (mc *MessageController) GetMessageList(w http.ResponseWriter, r *http.Request) {
	rows, err := mc.DB.Query("SELECT id, content, created_at FROM messages")
	if err != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messageList []models.Message
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.ID, &message.Content, &message.CreatedAt)
		if err != nil {
			http.Error(w, "Failed to scan message", http.StatusInternalServerError)
			return
		}
		messageList = append(messageList, message)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messageList)
}
