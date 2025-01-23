package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux" // Make sure this package is imported and used
)

// Message represents the structure of the incoming message
type Message struct {
	Content string `json:"content"`
}

// handleMessage is the handler function for POST requests to /messages
func handleMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message

	// Decode the incoming JSON payload into the Message struct
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		log.Printf("Error decoding message: %v", err)
		return
	}

	// Print the received message to the server log
	log.Printf("Received message: %s", msg.Content)

	// Send a response back
	response := fmt.Sprintf("Message received: %s", msg.Content)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Define the route for the POST request
	r.HandleFunc("/messages", handleMessage).Methods("POST")

	// Start the server on port 8080
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
