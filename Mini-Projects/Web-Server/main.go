package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func getCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05") 
}

// Handler for form submission
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() err: %v", err), http.StatusBadRequest)
		return
	}

	// Get current time
	currentTime := getCurrentTime()

	// Log request time
	log.Printf("Form submission received at %s", currentTime)

	name := r.FormValue("name")
	address := r.FormValue("address")

	// Send response with time
	response := fmt.Sprintf("Post request successful at %s\nName: %s\nAddress: %s", currentTime, name, address)
	fmt.Fprintf(w,"%s", response)
}

// Handler for "/hello" route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	// Get current time
	currentTime := getCurrentTime()

	// Log request time
	log.Printf("Hello request received at %s", currentTime)

	// Send response
	fmt.Fprintf(w, "Hello! Request received at: %s", currentTime)
}

func main() {
	// Serve static files from "static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Register handlers
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Start server
	port := ":8000"
	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
