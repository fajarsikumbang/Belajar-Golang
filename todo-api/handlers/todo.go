package handlers

import (
	"encoding/json"
	"net/http"
	"todo-api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// CreateTodo membuat to-do baru
func CreateTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo models.Todo
		_ = json.NewDecoder(r.Body).Decode(&todo)
		if err := db.Create(&todo).Error; err != nil {
			http.Error(w, "Gagal membuat to-do", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)
	}
}

// GetTodos mengembalikan semua to-do
func GetTodos(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todos []models.Todo
		if err := db.Find(&todos).Error; err != nil {
			http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	}
}

// GetTodo mengembalikan satu to-do berdasarkan ID
func GetTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var todo models.Todo
		if err := db.First(&todo, params["id"]).Error; err != nil {
			http.Error(w, "To-do tidak ditemukan", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todo)
	}
}

// UpdateTodo memperbarui to-do berdasarkan ID
func UpdateTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var todo models.Todo
		if err := db.First(&todo, params["id"]).Error; err != nil {
			http.Error(w, "To-do tidak ditemukan", http.StatusNotFound)
			return
		}
		_ = json.NewDecoder(r.Body).Decode(&todo)
		if err := db.Save(&todo).Error; err != nil {
			http.Error(w, "Gagal memperbarui to-do", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todo)
	}
}

// DeleteTodo menghapus to-do berdasarkan ID
func DeleteTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var todo models.Todo
		if err := db.First(&todo, params["id"]).Error; err != nil {
			http.Error(w, "To-do tidak ditemukan", http.StatusNotFound)
			return
		}
		if err := db.Delete(&todo).Error; err != nil {
			http.Error(w, "Gagal menghapus to-do", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
