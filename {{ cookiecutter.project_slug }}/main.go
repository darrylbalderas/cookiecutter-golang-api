package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ResponseData struct represents the structure of the response data
type ResponseData struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Set the content type header to indicate the type of data being returned
	w.Header().Set("Content-Type", "application/json")

	// Create a ResponseData object with required data
	response := ResponseData{
		Message: "Hello world",
	}

	// Marshal the response data into JSON format
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response to the http.ResponseWriter
	w.Write(jsonResponse)
}

func main() {
	port := "8000"
	http.HandleFunc("/", handler)
	log.Printf("Server is running on http://0.0.0.0:%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
