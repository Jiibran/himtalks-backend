package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"himtalks-backend/models"
	"himtalks-backend/ws"
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

	// Validasi kategori
	if message.Category != "kritik" && message.Category != "saran" {
		http.Error(w, "Category must be 'kritik' or 'saran'", http.StatusBadRequest)
		return
	}

	// Cek blacklist
	blacklisted, err := models.IsBlacklisted(mc.DB, message.Content)
	if err != nil {
		http.Error(w, "Error checking blacklist", http.StatusInternalServerError)
		return
	}
	if blacklisted {
		http.Error(w, "Message contains blacklisted word", http.StatusForbidden)
		return
	}

	// Simpan pesan ke database
	query := `INSERT INTO messages (content, sender_name, recipient_name, category) 
              VALUES ($1, $2, $3, $4) 
              RETURNING id, created_at`

	err = mc.DB.QueryRow(
		query,
		message.Content,
		message.SenderName,
		message.RecipientName,
		message.Category,
	).Scan(&message.ID, &message.CreatedAt)

	if err != nil {
		http.Error(w, "Failed to save message: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim pesan ke WebSocket
	ws.BroadcastMessage(ws.Message{
		Type: "message",
		Data: message,
	})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}

// GetMessageList mengembalikan daftar pesan
func (mc *MessageController) GetMessageList(w http.ResponseWriter, r *http.Request) {
	rows, err := mc.DB.Query("SELECT id, content, sender_name, recipient_name, category, created_at FROM messages")
	if err != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messageList []models.Message
	for rows.Next() {
		var message models.Message
		err := rows.Scan(
			&message.ID,
			&message.Content,
			&message.SenderName,
			&message.RecipientName,
			&message.Category,
			&message.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Failed to scan message", http.StatusInternalServerError)
			return
		}
		messageList = append(messageList, message)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messageList)
}

// DeleteMessage menghapus pesan berdasarkan ID
func (mc *MessageController) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var data struct{ ID int }
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	_, err := mc.DB.Exec("DELETE FROM messages WHERE id=$1", data.ID)
	if err != nil {
		http.Error(w, "Failed to delete message", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	// Kirim pesan ke WebSocket
	msg := ws.Message{
		Type: "delete",
		Data: data.ID,
	}
	ws.BroadcastMessage(msg)
}
