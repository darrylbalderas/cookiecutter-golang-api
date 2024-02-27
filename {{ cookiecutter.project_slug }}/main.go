package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ResponseData struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := ResponseData{
		Message: "Hello world",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func main() {
	port := "{{ cookiecutter.project_port }}"
	http.HandleFunc("/", handler)
	log.Printf("Server is running on http://0.0.0.0:%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
