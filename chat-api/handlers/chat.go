package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"chat-api/models"

	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

func ChatHandler(db *gorm.DB) {
	// Goroutine untuk menangani pesan broadcast
	go func() {
		for {
			msg := <-broadcast
			// Simpan pesan ke database
			db.Create(&msg)

			// Kirim pesan ke semua klien
			for client := range clients {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("Error: %v", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}()
}

func WebSocketEndpoint(w http.ResponseWriter, r *http.Request) {
	// Upgrade koneksi ke WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Tambahkan klien ke daftar
	clients[conn] = true

	for {
		var msg models.Message
		// Baca pesan dari klien
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error: %v", err)
			delete(clients, conn)
			break
		}
		msg.Timestamp = time.Now()
		broadcast <- msg
	}
}

func GetMessages(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var messages []models.Message
		if err := db.Find(&messages).Error; err != nil {
			http.Error(w, "Gagal mengambil pesan", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)
	}
}
