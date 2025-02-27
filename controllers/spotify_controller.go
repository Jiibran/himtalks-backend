package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"himtalks-backend/services"
)

type SpotifyController struct {
	SpotifyService *services.SpotifyService
}

func NewSpotifyController() *SpotifyController {
	spotifyService := services.NewSpotifyService(
		os.Getenv("SPOTIFY_CLIENT_ID"),
		os.Getenv("SPOTIFY_CLIENT_SECRET"),
	)
	return &SpotifyController{SpotifyService: spotifyService}
}

func (sc *SpotifyController) SearchTracks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			limit = 10
		}
	}

	tracks, err := sc.SpotifyService.SearchTracks(query, limit)
	if err != nil {
		http.Error(w, "Failed to search tracks: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tracks)
}

func (sc *SpotifyController) GetTrack(w http.ResponseWriter, r *http.Request) {
	trackID := r.URL.Query().Get("id")
	if trackID == "" {
		http.Error(w, "Track ID is required", http.StatusBadRequest)
		return
	}

	track, err := sc.SpotifyService.GetTrack(trackID)
	if err != nil {
		http.Error(w, "Failed to get track: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(track)
}
