package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"himtalks-backend/models"
)

type AdminHandler struct {
	DB *sql.DB
}

// Menambahkan admin baru
func (ah *AdminHandler) AddAdmin(w http.ResponseWriter, r *http.Request) {
	// Pastikan sudah dicek IsAdmin di middleware
	var data struct{ Email string }
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	if data.Email == "" {
		http.Error(w, "Email required", http.StatusBadRequest)
		return
	}
	err := models.InsertAdmin(ah.DB, data.Email)
	if err != nil {
		http.Error(w, "Failed to add admin", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Ubah batas hari untuk songfess
func (ah *AdminHandler) UpdateSongfessDays(w http.ResponseWriter, r *http.Request) {
	var data struct{ Days string }
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	if err := models.SetConfig(ah.DB, "songfess_days", data.Days); err != nil {
		http.Error(w, "Failed to update config", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Tambah kata blacklist
func (ah *AdminHandler) AddBlacklistWord(w http.ResponseWriter, r *http.Request) {
	var data struct{ Word string }
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	if data.Word == "" {
		http.Error(w, "Word required", http.StatusBadRequest)
		return
	}
	err := models.InsertBlacklistWord(ah.DB, data.Word)
	if err != nil {
		http.Error(w, "Failed to add word", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
