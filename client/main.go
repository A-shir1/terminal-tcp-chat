package main

import (
	"log"
)

func main() {
	username := getUsername()
	client := NewClient(username)
	err := client.ConnectToServer(":3000")
	if err != nil {
		log.Fatalf("Unable to connect to server: %v", err)
	}
}
