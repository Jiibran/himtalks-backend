package routes

import (
	"database/sql"

	"himtalks-backend/controllers"
	"himtalks-backend/ws"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Endpoint untuk pesan anonim dan songfess
	messageController := &controllers.MessageController{DB: db}
	songfessController := &controllers.SongfessController{DB: db}
	r.HandleFunc("/message", messageController.SendMessage).Methods("POST")
	r.HandleFunc("/songfess", songfessController.SendSongfess).Methods("POST")

	// Endpoint untuk WebSocket
	r.HandleFunc("/ws", ws.HandleConnections)

	return r
}
