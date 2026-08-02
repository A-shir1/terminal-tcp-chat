package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func (c *Client) WriteLoop() {
	reader := bufio.NewReader(os.Stdin)
	encoder := json.NewEncoder(c.conn)

	for {
		select {
		case <-c.quitch:
			return
		default:
		}

		fmt.Printf("%s>", c.username)
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("Write loop input error: %v", err)
			}
			c.Stop()
			return
		}

		line = strings.TrimRight(line, "\r\n")
		if strings.TrimSpace(line) == "" {
			continue
		}

		msg := Message{
			From:    c.username,
			Payload: line,
		}
		if err := encoder.Encode(msg); err != nil {
			log.Printf("Write error: %v", err)
			c.Stop()
			return
		}

		if line == "/exit" {
			c.Stop()
			return
		}
	}
}

func (c *Client) ReadLoop() {
	decoder := json.NewDecoder(c.conn)
	for {
		select {
		case <-c.quitch:
			return
		default:
		}
		var msg Message
		if err := decoder.Decode(&msg); err != nil {
			log.Printf("Read error HERE: %v", err)
			close(c.quitch)
			return
		}
		if msg.From == "server" {
			fmt.Printf("\x1b[3m %s\n \x1b[0m", msg.Payload)
			continue
		}
		fmt.Printf("%s: %s\n", msg.From, string(msg.Payload))
	}
}
