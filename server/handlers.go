package main

import (
	"fmt"
	"log"
	"net"
)

func (s *Server) AcceptConnectionLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Printf("Error Accepting connection: %v\n", err)
			continue
		}

		go s.HandleConnection(conn)
	}
}

func (s *Server) HandleConnection(conn net.Conn) {
	defer conn.Close()
	client := s.RegisterConnection(conn)

	// conn.Write([]byte("Connected to the tcp server!\n"))

	s.ReadIncomingMessageLoop(client)

	log.Printf("Connection with %v closed\n", conn.RemoteAddr())
}

func (s *Server) RegisterConnection(conn net.Conn) *Client {
	username, err := GetUsername(conn)
	if err != nil {
		log.Printf("Unable to read username: %v", err)
		username = "Unknown User"
	}
	newClient := &Client{
		conn:              conn,
		addr:              conn.RemoteAddr().String(),
		username:          username,
		connectedToServer: true,
	}
	s.clients = append(s.clients, newClient)
	fmt.Printf("New connection with %s (%s) registered!\n", newClient.username, newClient.addr)
	msg := username + " connected!"
	s.BroadcastServerMessage(msg)
	fmt.Println(newClient)
	return newClient
}
