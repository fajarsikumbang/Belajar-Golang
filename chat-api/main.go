package main

import (
	"fmt"
	"log"
	"net/http"

	"chat-api/handlers"
	"chat-api/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	// Inisialisasi database
	var err error
	db, err = gorm.Open("sqlite3", "./chat.db")
	if err != nil {
		log.Fatal("Tidak dapat membuka database: ", err)
	}
	// Auto migrate untuk model Message
	db.AutoMigrate(&models.Message{})
}

func main() {
	r := mux.NewRouter()

	// WebSocket Endpoint
	handlers.ChatHandler(db)
	r.HandleFunc("/ws", handlers.WebSocketEndpoint)

	// REST API untuk riwayat pesan
	r.Handle("/messages", handlers.AuthMiddleware(http.HandlerFunc(handlers.GetMessages(db)))).Methods("GET")

	// Jalankan server
	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
