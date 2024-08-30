package main

import (
	"encoding/json"
	"log"
	"net/http"
	"talentpitch/cqrs"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/commands", handleCommand).Methods("POST")

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func handleCommand(w http.ResponseWriter, r *http.Request) {
	var command cqrs.Command
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&command); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	event, err := cqrs.BuildCommand(command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}
