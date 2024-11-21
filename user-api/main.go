package main

import (
    "log"
    "net/http"
    "user-api/handlers"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // Endpoint Routes
    r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

    // Start server
    log.Println("Server berjalan di http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
