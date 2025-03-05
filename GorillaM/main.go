package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var userCache = make(map[int]User)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var cacheMutex sync.RWMutex

func main() {
	r := mux.NewRouter()

	// Defining the routes and their handlers
	r.HandleFunc("/", handleRoot)
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", getUser).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", deleteUser).Methods("DELETE")

	// Start the server
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Extract the id from the URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := userCache[id]; !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	cacheMutex.Lock()
	delete(userCache, id)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cacheMutex.RLock()
	user, ok := userCache[id]
	cacheMutex.RUnlock()
	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// Decode the user from the request body
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	userCache[len(userCache)+1] = user
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	//Extract the id from the URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Decode the updated user from the request body
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if updatedUser.Name == "" {
		http.Error(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}
	if updatedUser.Age <= 0 {
		http.Error(w, "Age must be greater than zero", http.StatusBadRequest)
		return
	}

	// Update the user in the cache
	cacheMutex.Lock()
	existingUser, exists := userCache[id]
	if !exists {
		cacheMutex.Unlock()
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Update user details
	existingUser.Name = updatedUser.Name
	existingUser.Age = updatedUser.Age
	userCache[id] = existingUser
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusOK)
}
