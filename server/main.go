package main

import (
	"fmt"
	"log"
)

func main() {
	server := NewServer(":3000")

	fmt.Println("Starting Server...")
	err := server.Start()
	if err != nil {
		log.Fatalf("Error staring the server: %v\n", err)
	}
}
