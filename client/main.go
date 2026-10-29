package main

import (
	"fmt"
	"log"
)

func main() {
	username := getUsername()
	client := NewClient(username)
	fmt.Println("Connecting to server")
	err := client.ConnectToServer(":3000")
	if err != nil {
		log.Fatalf("Unable to connect to server: %v", err)
	}
}
