package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/my-gorilla-mux-app/internal/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/friends", handlers.GetFriends).Methods("GET")
	router.HandleFunc("/friends/{name}", handlers.GetFriendByName).Methods("GET")
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
