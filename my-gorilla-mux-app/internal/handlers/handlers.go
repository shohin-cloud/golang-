package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friends)
}

func GetFriendByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := params["name"]
	for _, friend := range friends {
		if friend.Name == name {
			json.NewEncoder(w).Encode(friend)
			return
		}
	}
	http.Error(w, "Friend not found", http.StatusNotFound)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("My Gorilla Mux App - Author: Your Name"))
}
