package main

import (
	"log"
)

func main() {
	server := NewServer(":3000")

	err := server.Start()
	if err != nil {
		log.Fatalf("Error staring the server: %v\n", err)
	}
}
