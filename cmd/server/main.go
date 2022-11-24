package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"mailganer-task/internal/db"
	"mailganer-task/internal/routes"
)

// Start a hello world api on port 8080
func main() {

	// Test connection to the database
	db := db.Connect()

	// Create a new router
	r := mux.NewRouter()

	_ = db

	// Root route
	r.HandleFunc("/", routes.Root).Methods("GET")

	// Get all users route
	r.HandleFunc("/users", routes.GetAllUsers).Methods("GET")

	// Start server on port 8080
	http.ListenAndServe(":8080", r)
}
