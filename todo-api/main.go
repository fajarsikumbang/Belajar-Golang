package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-api/handlers"  // Mengimpor handlers
	"todo-api/models"    // Mengimpor models

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB global variable
var db *gorm.DB

func init() {
	// Inisialisasi koneksi database
	var err error
	// Menggunakan SQLite sebagai database
	db, err = gorm.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal("Tidak dapat membuka database: ", err)
	}
	// Auto migrate untuk model Todo
	db.AutoMigrate(&models.Todo{})
}

func main() {
	// Membuat router menggunakan mux
	r := mux.NewRouter()

	// Menetapkan route
	r.HandleFunc("/todos", handlers.CreateTodo(db)).Methods("POST")
	r.HandleFunc("/todos", handlers.GetTodos(db)).Methods("GET")
	r.HandleFunc("/todos/{id}", handlers.GetTodo(db)).Methods("GET")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo(db)).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo(db)).Methods("DELETE")

	// Mulai server
	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
