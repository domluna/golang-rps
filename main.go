package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Notification represents the structure of a notification
type Notification struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// Queue to store notifications
var notificationQueue = make(chan Notification, 100)
var wg sync.WaitGroup

// Handler for POST requests to add notifications
func notificationHandler(w http.ResponseWriter, r *http.Request) {
	// simulate average of DB/cache call for user info
	time.Sleep(100 * time.Millisecond)

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var notification Notification
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	if notification.Type != "email" && notification.Type != "sms" && notification.Type != "push" {
		http.Error(w, "Invalid notification type", http.StatusBadRequest)
		return
	}

	// Add the notification to the queue
	notificationQueue <- notification

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Notification received")
}

// Worker to process notifications from the queue
func notificationWorker() {
	defer wg.Done()
	for notification := range notificationQueue {
		// Process the notification (for now, just print it)
		fmt.Printf("Processing notification: Type=%s, Content=%s\n", notification.Type, notification.Content)
	}
}

func main() {
	// Start the notification worker
	wg.Add(1)
	go notificationWorker()

	// Set up the HTTP server
	http.HandleFunc("/notify", notificationHandler)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// Close the queue and wait for the worker to finish
	close(notificationQueue)
	wg.Wait()
}
