package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

// Book struct
type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

// Mock database
var books = []Book{
    {ID: 1, Title: "1984", Author: "George Orwell"},
    {ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
}

// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

// Get a single book
func GetBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for _, book := range books {
        if book.ID == id {
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

// Create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = len(books) + 1
    books = append(books, book)
    json.NewEncoder(w).Encode(book)
}

// Update a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for index, book := range books {
        if book.ID == id {
            books = append(books[:index], books[index+1:]...)
            var updatedBook Book
            _ = json.NewDecoder(r.Body).Decode(&updatedBook)
            updatedBook.ID = id
            books = append(books, updatedBook)
            json.NewEncoder(w).Encode(updatedBook)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

// Delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for index, book := range books {
        if book.ID == id {
            books = append(books[:index], books[index+1:]...)
            json.NewEncoder(w).Encode(books)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}
