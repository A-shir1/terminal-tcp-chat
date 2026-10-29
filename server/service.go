package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
)

func (s *Server) PrintClients() {
	for _, client := range s.clients {
		fmt.Printf("Username: %s | Address: %v\n", client.username, client.addr)
	}
}

func (s *Server) GetMessageFrom(msg Message) string {
	for _, client := range s.clients {
		if client.addr == msg.From {
			return client.username
		}
	}
	return "UNKNOWN USER"
}

func (s *Server) ReadIncomingMessageLoop(client *Client) {
	decoder := json.NewDecoder(client.conn)
outer:
	for {
		var msg Message
		// 	~~~~~~	THIS IS MOST LIKELY NOT THE RIGHT WAY TO HANDLE THIS ERROR. FIX IT	~~~~~~~
		if err := decoder.Decode(&msg); err != nil {
			log.Printf("Read Error while trying to decode the incoming message: %v", err)
			// close(s.quitch)
			return
		}

		switch string(msg.Payload) {
		case "/exit":
			msg := client.username + " disconnected"
			s.BroadcastServerMessage(msg)
			break outer
		case "/showUsers":
			s.PrintClients()
		default:
			s.BroadcastMessage(msg)
		}
	}
}

func (s *Server) BroadcastMessage(msg Message) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Printf("JSON Marshal error: %v\n", err)
		return
	}

	// Add newline for protocol clarity
	jsonMsg = append(jsonMsg, '\n')

	for _, client := range s.clients {
		// msg.Payload += "\n"
		_, err := client.conn.Write(jsonMsg)
		if err != nil {
			log.Printf("Write error to %v: %v\n", client.addr, err)
			continue
		}
	}
}

func (s *Server) BroadcastServerMessage(line string) {
	msg := Message{
		From:    "server",
		Payload: line,
	}
	s.BroadcastMessage(msg)
}

func GetUsername(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	username = strings.TrimSpace(username)

	return username, nil
}
