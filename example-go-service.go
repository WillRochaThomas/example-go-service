package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type health struct {
	Status string `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	diagnostics := health{"UP"}

	json, err := json.Marshal(diagnostics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	println("Server running on port", 8080)

	http.HandleFunc("/health", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
