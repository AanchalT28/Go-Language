package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var (
	items = make(map[string]Item)
	mutex = &sync.Mutex{}
)

func main() {
	http.HandleFunc("/items", handleItems) //routes http requests with path /items to handleItems func
	http.HandleFunc("/items/", handleItem) //routes http requests with path /items/{id} to handleItem func
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleItems(w http.ResponseWriter, r *http.Request) { //handler for /items
	switch r.Method {
	case http.MethodGet:
		getAllItems(w, r)
	case http.MethodPost:
		createItems(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleItem(w http.ResponseWriter, r *http.Request) { //handler for /items/{id}
	switch r.Method {
	case http.MethodGet:
		getItem(w, r)
	case http.MethodPut:
		updateItem(w, r)
	case http.MethodDelete:
		deleteItem(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	defer mutex.Unlock()
	log.Println("Current items:", items)
	json.NewEncoder(w).Encode(items)
}

func createItems(w http.ResponseWriter, r *http.Request) {
	var newItems []Item
	if err := json.NewDecoder(r.Body).Decode(&newItems); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mutex.Lock()
	for _, item := range newItems {
		if _, exists := items[item.ID]; exists {
			http.Error(w, "Item already exists", http.StatusConflict)
			mutex.Unlock()
			return
		}
		items[item.ID] = item
	}
	mutex.Unlock()
	log.Println("Items added:", newItems)
	w.WriteHeader(http.StatusCreated)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/items/"):]
	item, ok := items[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/items/"):]
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mutex.Lock()
	items[id] = item
	mutex.Unlock()
	w.WriteHeader(http.StatusOK)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/items/"):]
	mutex.Lock()
	delete(items, id)
	mutex.Unlock()
	w.WriteHeader(http.StatusNoContent)
}
