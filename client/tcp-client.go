package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Client struct {
	conn     net.Conn
	username string
	quitch   chan struct{}
	stopOnce sync.Once
}

func NewClient(username string) *Client {
	return &Client{
		username: username,
		quitch:   make(chan struct{}),
	}
}

func (c *Client) Stop() {
	c.stopOnce.Do(func() {
		close(c.quitch)
	})
}

func (c *Client) ConnectToServer(serverAddr string) error {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	c.conn = conn

	// 	SENDS THE USERNAME DATA TO THE SERVER
	if _, err := fmt.Fprintf(conn, "%s\n", c.username); err != nil {
		return err
	}

	fmt.Println("CONNECTED TO SERVER!!")

	buf := make([]byte, 2048)
	if _, err := c.conn.Read(buf); err != nil {
		log.Printf("READ ERROR DURING WELCOME MESSAGE")
	}
	fmt.Println(string(buf))

	go c.ReadLoop()
	go c.WriteLoop()

	<-c.quitch
	return nil
}
