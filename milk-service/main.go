package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func MilkHandler(w http.ResponseWriter, r *http.Request) {
	milkType := os.Getenv("MILK_TYPE")
	if milkType == "" {
		http.Error(w, "MILK_TYPE is not set", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("%s", milkType)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(response))
}

func main() {

	http.HandleFunc("/milk", MilkHandler)

	port := os.Getenv("PORT")
	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
