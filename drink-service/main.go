package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func DrinkHandler(w http.ResponseWriter, r *http.Request) {
	drinkType := os.Getenv("DRINK_TYPE")
	if drinkType == "" {
		http.Error(w, "Drink_TYPE is not set", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("%s", drinkType)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(response))
}

func main() {

	http.HandleFunc("/", DrinkHandler)

	port := os.Getenv("PORT")
	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
