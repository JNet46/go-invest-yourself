package main

import (
	"encoding/json"
	"go-invest-yourself/db"
	"go-invest-yourself/handlers"
	"go-invest-yourself/models"
	"net/http"
)

func main() {
	// Create some mock users
	db.AddUser(&models.User{ID: "1", Name: "Alice", Balance: 100.0})
	db.AddUser(&models.User{ID: "2", Name: "Bob", Balance: 50.0})

	http.HandleFunc("/pay", handlers.MakePayment)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(db.Users)
	})

	http.ListenAndServe(":8080", nil)
}
