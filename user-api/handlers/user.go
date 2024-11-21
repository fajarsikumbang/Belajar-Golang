package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

// User struct
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Username string `json:"username"`
}

// Mock database
var users = []User{
    {ID: 1, Name: "John Doe", Email: "john.doe@example.com", Username: "johndoe"},
    {ID: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Username: "janesmith"},
}

// Get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// Get a single user
func GetUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for _, user := range users {
        if user.ID == id {
            json.NewEncoder(w).Encode(user)
            return
        }
    }
    http.Error(w, "User not found", http.StatusNotFound)
}

// Create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user User
    _ = json.NewDecoder(r.Body).Decode(&user)
    user.ID = len(users) + 1
    users = append(users, user)
    json.NewEncoder(w).Encode(user)
}

// Update a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for index, user := range users {
        if user.ID == id {
            users = append(users[:index], users[index+1:]...)
            var updatedUser User
            _ = json.NewDecoder(r.Body).Decode(&updatedUser)
            updatedUser.ID = id
            users = append(users, updatedUser)
            json.NewEncoder(w).Encode(updatedUser)
            return
        }
    }
    http.Error(w, "User not found", http.StatusNotFound)
}

// Delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for index, user := range users {
        if user.ID == id {
            users = append(users[:index], users[index+1:]...)
            json.NewEncoder(w).Encode(users)
            return
        }
    }
    http.Error(w, "User not found", http.StatusNotFound)
}
