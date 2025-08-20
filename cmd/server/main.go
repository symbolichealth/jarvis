package main

import (
	"log"

	jarvis "github.com/symbolichealth/jarvis/backend"
)

func main() {
	log.Println("Starting Jarvis web server...")

	server := jarvis.NewServer()
	if err := server.StartServer("7070"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
