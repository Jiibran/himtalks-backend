package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Log headers untuk debugging
	log.Println("Headers:", r.Header)

	// Upgrade koneksi HTTP ke WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("WebSocket upgrade error: %v", err)
		return
	}
	defer ws.Close()

	log.Println("Client connected via WebSocket")

	// Loop untuk membaca dan menulis pesan
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}

		log.Printf("Received message: %s", string(p))

		// Kirim pesan balik ke client
		if err := ws.WriteMessage(messageType, p); err != nil {
			log.Printf("WebSocket write error: %v", err)
			break
		}
	}
}
