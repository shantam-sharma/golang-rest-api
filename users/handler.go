package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/*
|| Module Docs ||
- Handler only manages http requests.
- Handles HTTP requests and responses.
- eads request data, calls services, returns JSON/status codes.
*/

var store = NewStore()
var service = NewService(store)

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Pong")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	user := service.CreateUser(req.Name, req.Email)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	// 2 variables userID and err
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	//check if user exist.
	user, exists := service.GetUser(userID)
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	//Client ko batata hai ki response ka data JSON format me aa raha hai.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	_, exists := service.GetUser(userID)
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	service.DeleteUser(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "user deleted successfully",
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	_, exists := service.GetUser(userID)
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	var req User

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	updatedUser := service.UpdateUser(userID, req.Name, req.Email)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(updatedUser)
}
