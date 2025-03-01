package main

import (
	"fmt"
	"inno-lb-sidecar/internal/sidecar"
	"log"
	"net/http"
	"os"
)

func main() {
	sideCar := sidecar.NewSidecar()

	sideCar.Register()
	go sideCar.Ping()

	http.HandleFunc("/info", sideCar.InfoHandler)
	http.HandleFunc("/", sideCar.ProxyHandler)

	port := os.Getenv("PORT")
	fmt.Printf("Sidecar running on port %s, proxying to master at %s\n", port, os.Getenv("MASTER_URL"))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
