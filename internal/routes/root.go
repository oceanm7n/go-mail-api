package routes

// Root is the root route, returns a simple message

// Path: internal\routes\root.go
import (
	"encoding/json"
	"mailganer-task/internal/models"
	"net/http"
)

// Root is the root route, returns a simple message
func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

// Get all users route
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(users)
}
