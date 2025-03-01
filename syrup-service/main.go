package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func SyrupHandler(w http.ResponseWriter, r *http.Request) {
	syrupType := os.Getenv("SYRUP_TYPE")
	if syrupType == "" {
		http.Error(w, "SYRUP_TYPE is not set", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("%s", syrupType)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(response))
}

func main() {

	http.HandleFunc("/", SyrupHandler)

	port := os.Getenv("PORT")
	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
